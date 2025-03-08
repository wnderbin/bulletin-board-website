<h1 align="center"> Bulletin board website </h1>

![image](https://github.com/user-attachments/assets/eed17f44-ff9f-4da7-a8ba-1f4c9c6638be)

<h4 align="center"> Bulletin board website with CRUD database and statistical files  </h4>
<h3 align="center"> Before using, I advise you to read the license </h3>

## ABOUT

**Language:** Go 1.24.0

**Tested on:** Linux

**Dependencies:** Sqlite3 driver (go get github.com/mattn/go-sqlite3)

**Author:** wnderbin

## INSTALLING AND LAUNCH

```
git clone https://github.com/wnderbin/bulletin-board-website

( You may need to install the sqlite3 driver )
( go get github.com/mattn/go-sqlite3 )

cd bulletin-board-website/cmd/bulletin-board
go run main.go

---

(You can specify your port, but then the navigation on the site will not work correctly.)
(To do this, you will need to correct port 8080 to yours in the files.)
(Files: delete_bulletin.html; main_page.html; not_found.html; send_form.html; update_bulletin)

go run main.go 8081
```

### PREREQUISITES
* Install sqlite3 driver with go get ...
