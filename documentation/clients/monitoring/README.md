# Monitoring

Honeybadger is the centralised system of choice for exception monitoring. The
client can be imported with:

```ruby
require './src/clients/monitoring'
```

and invoked with:

```ruby
Clients::Monitoring.notify('Hello, world')
```

Honeybadger requires the `HONEYBADGER_API_KEY` env variable to work.
