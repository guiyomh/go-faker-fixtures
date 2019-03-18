package cmd

import (
	"errors"
	"fmt"
	"os"

	internalcontract "github.com/guiyomh/charlatan/internal/contracts"
	treecontracts "github.com/guiyomh/charlatan/pkg/tree/contracts"

	"github.com/azer/logger"
	"github.com/guiyomh/charlatan/internal/pkg/db"
	"github.com/guiyomh/charlatan/internal/pkg/generator"
	"github.com/guiyomh/charlatan/internal/pkg/reader"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(loadCmd)
}

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load fixtures from the path",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires one arg")
		}
		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			return fmt.Errorf("the directory '%s' not existing", args[0])
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.New("cmd")
		timer := log.Timer()
		fixturePath := args[0]
		reader := reader.NewFixtureReader()
		data, err := reader.Read(fixturePath)

		if err != nil {
			log.Error(err.Error())
			panic(err)
		}

		//generator := generator.NewGenerator()
		rows, err := generator.NewGenerator().GenerateRecords(data)
		if err != nil {
			log.Error(err.Error())
			panic(err)
		}
		tree, err := generator.BuildTree(rows)
		if err != nil {
			log.Error(err.Error())
			panic(err)
		}
		dbManagerFactory := db.DbManagerFactory{}
		manager, err := dbManagerFactory.NewDbManager("mysql", DbHost, DbPort, DbUser, DbPass)
		if err != nil {
			panic(err)
		}
		tree.Walk(func(node treecontracts.Node) {
			original, _ := node.(internalcontract.Row)
			if original.HasDependencies() {
				for field, relation := range original.DependencyReference() {
					target, _ := tree.Find(relation.RecordName()).(internalcontract.Row)
					if relation.FieldName() != "" {
						original.Fields()[field] = target.Fields()[relation.RecordName()]
					} else {
						original.Fields()[field] = target.Pk
					}
				}
			}
			sql, params, err := manager.BuildInsertSQL(original.Schema(), original.TableName(), original.Fields())
			if err != nil {
				panic(err)
			}
			result, err := manager.Exec(sql, params)
			if err != nil {
				panic(err)
			}
			lastInsertID, err := result.LastInsertId()
			if err != nil {
				panic(err)
			}
			original.SetPk(lastInsertID)
		}, true)

		log.Info(fmt.Sprintf("Nb rows : %d", len(rows)))

		timer.End("Insert record in database")
	},
}
