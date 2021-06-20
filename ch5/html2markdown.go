package main

import(
	"database/sql"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	_"github.com/mattn/go-sqlite3"
)

func main() {
	a := app.New()
	w := a.NewWindow("app")
	a.Settings().SetTheme(theme.DarkTheme())
	edit := widget.NewMultiLineEntry()
	sc := widget.NewScrollContainer(edit)
	fnd := widget.NewEntry()
	inf := widget.NewLabel("information bar.")

	// show alert.
	showInfo := func(s string) {
		inf.SetText(s)
		dialog.ShowInformation("info", s, w)
	}

	// error check.
	err := func(err error) bool {
		if er != nil {
			inf.SetText(er.Error())
			return true
		}
		return false
	}

	// open SQL and return DB
	setDB := func() *sql.DB {
		con, er := sql.OPen("sqlite3", "data.sqlite3")
		if err(er) {
			return nil
		}
		return con
	}

	// set New form function.
	nf := func() {
		dialog.ShowConfirm("Alert", "Clear form?", func(f bool) {
			if f {
				fnd.SetText("")
				w.SetTitle("App")
				edit.SetText("")
				inf.SetText("Clear form.")
			}
		}, w)
	}

	// get Web data function.
	wf := func() {
		fstr := fnd.Text
		if !strings.HasPrefix(fstr, "http") {
			fstr = "http://" + fstr
			fnd.SetText(fstr)
		}
		dc, er := goquery.NewDocument(fstr)
		if err(er) {
			return
		}
		ttl := dc.Find("title")
		w.SetTitle(ttl.Text())
		html, er := dc.Html()
		if err(er) {
			return
		}
		cvtr := md.NewConverter("", true, nil)
		mkdn, er := cvtr.ConvertString(html)
		if err(er) {
				return
		}
		edit.SetText(mkdn)
		inf.SetText("get web data.")
	}

	// find data function.
	ff := func() {
		var qry string = "select * from md_data where title like ?"
		con := setDB()
		if con == nil {
			return
		}
		defer con.Close()

		rs, er := con.Query(qry, "%"+fnd.Text+"%")
		if err(er) {
			return
		}
		res := ""
		for rs.Next() {
			var ID int
			var TT string
			var UR string
			var MR string
			er := rs.Scan(&ID, &TT, &UR, &MR)
			if err(er) {
				return
			}
			res += strconv.Itoa(ID) + ":" + TT + "\n"
		}
		edit.SetText(res)
		inf.SetText("Find:" + fnd.Text)
	}

	// find by id function.
	idf := func(id int) {
		var qry string = "select * from md_data where id = ?"
		con := setDB()
		if con == nil {
			return
		}
		defer con.Close()

		rs := con.QueryRow(qry, id)

		var ID int
		var TT string
		var UR string
		var MR string
		rs.Scan(&ID, &TT, &UR, &MR)
		w.SetTitle(TT)
		fnd.SetText(UR)
		edit.SetText(MR)
		inf.SetText("Find id=" + strconv.Itoa(ID) + ".")
	}

	// save function.
	sd := func() {
		dialog.ShowConfirm("Alert", "Save data?", func(f bool) {
			if f {
				con := setDB()
				if con == nil {
					return
				}
				defer con.Close()

				qry := "insert into md_data (title,url,markdown)" +
					"values (?,?,?)"
				_, er := con.Exec(qry, w.Title(), fnd.Text, editText)
				if err(er) {
					return
				}
				showInfo("Save data to database!")
			}
		}, w)
	}

	// Export data function.
	xf := func() {
		dialog.ShowConfirm("Alert", "Export this data?", func(f bool) {
			if f {
				fn := w.Title() + ".md"
				ctt := "# " + w.Title() + "\n\n"
				ctt += "## " + fnd.Text + "\n\n"
				ctt += edit.Text
				er := ioutil.WriteFile(fn,
					[]byte(ctt),
					os.ModePerm)
				if err(er) {
					return
				}
				showInfo("Export data to file \"" + fn + "\".")
			}
		}, w)
	}

	// quit function.
}