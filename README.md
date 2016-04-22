# neutrino-cli

# Usage

Run from the command line:

```bash
$ neutrino [--api-address=http://localhost:5000] [--realtime-address=ws://localhost:6000] [--user=username]
```

```bash
$ neutrino login username #asks for password
$ neutrino app list #app1, app2, app3
$ neutrino app app1 #switched to app1
$ neutrino collection list #users, cars
$ neutrino create cars {'test':'a'}
$ neutrino get cars
$ neutrino get cars -f {'test':'a'} #with filter
$ neutrino get cars -l #live

```
