# Reference

## Creating fixtures

The most basic functionality of this library is to turn flat yaml files into map of fields.
You can define many maps of different table in one file as such:

```yaml
user:
    admin_1:
        first_name: 'William'
        last_name: 'Wallace'
    admin_2:
        first_name: 'Bob'
        last_name: 'The sponge'

group:
    group_1:
        name: admin
    group_2:
        name: reader
```

## Fixture Ranges

The first step is to let Charlatan create many copies of a map for you to remove duplication from the yaml file.

You can do that by defining a range in the fixture name:

```yaml
user:
    user_{1..10}:
        first_name: 'William'
        last_name: 'Wallace'
        email: 'william@example.org'
```
Now it will generate ten users, with IDs user_1 to user_10. Pretty good but we only have 10 williams with the same name, lastname and email which is not so fancy yet.

## Fixture Lists

You can also specify a list of values instead of a range:
```yaml
user:
    user_{william, bob}:
        first_name: '<Current()>'
        last_name: 'Wallace'
        email: '<Current()>@example.org'
```

The `<Current()>` function is a bit special as it can only be called in the context of a collection (list of values or a range).

In the case of a list of values like the example above, it will return for the first fixture user_william the value william, and bob for the fixture user_bob.

In the case of a range (e.g. user{1..10}), `<Current()>` will return 1 for user1, 2 for user2 etc.

To go further we the example above, we can just randomize data.