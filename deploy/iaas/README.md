# Notes

* Install PG 14 with [the official doc](https://www.postgresql.org/download/linux/debian/)
* With user postgres, run `ALTER USER postgres WITH PASSWORD 'secret';` to add password to the user
* add unit file to `/etc/systemd/system/demo.service`, then enable it