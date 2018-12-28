package main

import (
	"bufio"
	"bytes"
	"ee/schkola/person"
	"ee/schkola/shared"
	"ee/schkola/student"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/looplab/eventhorizon"
	"github.com/yvasiyarov/php_session_decoder/php_serialize"
)

var work = "/Users/ee/Google Drive/Bibelschule/Stephanus/Verwaltung/BSS-2017-2018"

var de = map[string]string{
	"scool_year":         "Klasse",
	"biblikum_1":         "Biblikum 1 geschrieben?",
	"biblikum_2":         "Biblikum 2 geschrieben?",
	"paided_last_years":  "Gebühr der Vorjahren bezahlt?",
	"first_name":         "Vorname",
	"last_name":          "Nachname",
	"birth_date":         "Geburtsdatum",
	"gender":             "Geschlecht",
	"phone_number":       "Telefon",
	"church":             "Gemeinde",
	"church_commitment":  "Gemeinde einverstanden?",
	"church_member":      "Mitglied welcher Gemeinde?",
	"church_services":    "Gemeindedienste",
	"church_responsible": "Pastor / Leitungskreis",
	"church_contact":     "Telefon von Pastor / Leitungskreis",
	"user_email":         "E-Mail",
	"address":            "Adresse",
	"plz":                "PLZ",
	"city":               "Ort",
	"job":                "Beruf",
	"education":          "Bildung",
	"marital_state":      "Familienstand",
	"spirit":             "Geistlicher_Werdegang______________________________________________________Warum_Bibelschule?",
}

func main() {
	parseJson()
}

func parseJson() {

	file, err := ioutil.ReadFile(fmt.Sprintf("%v/wp_usermeta.json", work))
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		panic(err)
	}

	var data []map[string]interface{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Printf("JSON unmarshal error: %v\n", err)
		panic(err)
	}

	keys := make(map[string]string)

	u := make(map[string]map[string]interface{})

	toDec, _ := regexp.Compile("^[a-z]\\:.*")

	for _, d := range data {
		user := d["user_id"].(string)
		key := d["meta_key"].(string)
		str := fmt.Sprintf("%v", d["meta_value"])

		userMap := u[user]
		if u[user] == nil {
			userMap = make(map[string]interface{})
			u[user] = userMap
		}

		if toDec.MatchString(str) {
			if dec, err := php_serialize.NewUnSerializer(str).Decode(); err == nil && dec != nil {
				fill(key, dec, userMap, keys)
			} else {
				println(dec, err)
			}
		} else {
			keys[key] = key
			userMap[key] = prepareString(str)
		}
	}

	var wpUsers []map[string]interface{}
	file, err = ioutil.ReadFile(fmt.Sprintf("%v/wp_users.json", work))
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		panic(err)
	}
	err = json.Unmarshal(file, &wpUsers)
	if err != nil {
		fmt.Printf("JSON unmarshal error: %v\n", err)
		panic(err)
	}

	for _, wpU := range wpUsers {
		id := wpU["ID"].(string)
		userMap := u[id]

		if userMap == nil {
			userMap = make(map[string]interface{})
			u[id] = userMap
		}

		for k, v := range wpU {
			keys[k] = k
			userMap[string(k)] = v
		}

	}
	//println(fmt.Sprintf("%v", keys))

	var allKeys []string
	for key, _ := range keys {
		allKeys = append(allKeys, key)
	}

	users := make(map[string]map[string]interface{})
	users1 := make([]string, 0)
	users2 := make([]string, 0)
	users3 := make([]string, 0)
	users4 := make([]string, 0)
	users5 := make([]string, 0)
	besucher := make([]string, 0)

	/*
		for _, v := range u {
			key := fmt.Sprintf("%v.%v.%v.%v", v["scool_year"], v["last_name"], v["first_name"], v["birth_date"])
			users[key] = v
			switch v["scool_year"] {
			case "1. Klasse":
				users1 = append(users1, key)
			case "2. Klasse (B1)":
				users2 = append(users2, key)
			case "3. Klasse (B1)":
				users3 = append(users3, key)
			case "4. Klasse (B1,B2)":
				users4 = append(users4, key)
			case "Zusatzjahr (B1,B2)":
				users5 = append(users5, key)
			case "Besucher":
				besucher = append(besucher, key)
			}
		}*/

	for _, v := range u {
		email, _ := v["user_email"]
		status, ok := v["account_status"]
		if email == "viktor.itermann@gmx.de" || (ok && status == "awaiting_admin_review") {
			key := fmt.Sprintf("%v.%v.%v.%v", v["wp_capabilities"], v["last_name"], v["first_name"], v["birth_date"])
			users[key] = v
			yearObj, ok := v["scool_year"]
			year := yearObj.(string)
			if ok {
				if strings.Contains(year, "1. Klasse") {
					users1 = append(users1, key)
				} else if strings.Contains(year, "2. Klasse") {
					users2 = append(users2, key)
				} else if strings.Contains(year, "3. Klasse") {
					users3 = append(users3, key)
				} else if strings.Contains(year, "4. Klasse") {
					users4 = append(users4, key)
				} else if strings.Contains(year, "Zusatzjahr") {
					users5 = append(users5, key)
				}
			} else if _, ok := v["guest"]; ok {
				besucher = append(besucher, key)
			}
		}
	}

	writeCsvEmailGroup(users1, users2, users3, users4, users5, besucher, users)
	//writeInsertInto(users1, users2, users3, users4, users5, users)
	//writeUpdateRole(users1, users2, users3, users4, users5, users)
	writeCsvData(users1, users2, users3, users4, users5, besucher, users)
	writeCsvEmailContacts(users1, users2, users3, users4, users5, besucher, users)
	writeHtmlReport(users1, users2, users3, users4, users5, besucher, allKeys, users)
	//restImport(users1, users2, users3, users4, users5, users)
}

func restImport(users1 []string, users2 []string, users3 []string,
	users4 []string, users5 []string, besucher []string, users map[string]map[string]interface{}) {

	client := http.Client{
		Timeout: time.Second * 10, // Maximum of 2 secs
	}

	restImportForGroup(users1, users, "2017 Klasse 1", &client)
	restImportForGroup(users2, users, "2016 1. Klasse", &client)
	restImportForGroup(users3, users, "2016 2. Klasse", &client)
	restImportForGroup(users4, users, "2016 3. Klasse", &client)
	restImportForGroup(users5, users, "2016 4. Klasse", &client)
}

func restImportForGroup(userKeys []string, users map[string]map[string]interface{}, group string, client *http.Client) {

	profilesUrl := "http://127.0.0.1:8080/person/profiles/"
	profiles := make(map[eventhorizon.UUID]*person.Profile)
	//applications := make(map[eventhorizon.UUID]*student.SchoolApplication)

	for _, userKey := range userKeys {
		user := users[userKey]
		profile := person.NewProfile()

		profile.Id = UUID(user["user_login"])
		profile.Gender, _ = person.Genders().ParseGenderGerman(str(user["gender"]), person.Genders().Unknown())
		profile.Name = shared.NewPersonName()
		profile.Name.First = str(user["first_name"])
		profile.Name.Last = str(user["last_name"])
		profile.BirthName = str(user["birthname"])
		profile.Birthday = timeValue(user["birth_date"])
		profile.Address = person.NewAddress()
		profile.Address.City = str(user["city"])
		profile.Address.Code = str(user["plz"])
		profile.Address.Country = "Deutschland"
		profile.Address.Street = str(user["address"])
		profile.Contact = person.NewContact()
		profile.Contact.Email = str(user["user_email"])
		profile.Contact.Phone = str(user["phone_number"])
		//profile.PhotoData = blob(str(user["photo"]))
		profile.Family = person.NewFamily()
		profile.Family.MaritalState, _ = person.MaritalStates().ParseMaritalStateGerman(str(user["marital_state"]),
			person.MaritalStates().Unknown())
		profile.Family.Partner, _ = person.PersonNameParse(str(user["partner"]))
		profile.Church = person.NewChurchInfo()
		profile.Church.Church = str(user["church"])
		profile.Church.Member = !strings.EqualFold(str(user["church_member"]), "Keine Mitgliedschaft")
		profile.Church.Services = str(user["church_services"])
		profile.Education = person.NewEducation()
		profile.Education.Other = str(user["education"])
		profile.Education.Profession = str(user["job"])

		//profile.Church.Association = str(user["church_member"])

		profiles[profile.Id] = profile

		/*

			type SchoolApplication struct {
				Profile *person.Profile `json:"profile" eh:"optional"`
				RecommendationOf *shared.PersonName `json:"recommendationOf" eh:"optional"`
				ChurchContactPerson *shared.PersonName `json:"churchContactPerson" eh:"optional"`
				ChurchContact *person.Contact `json:"churchContact" eh:"optional"`
				ChurchCommitment bool `json:"churchCommitment" eh:"optional"`
				SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
				Group string `json:"group" eh:"optional"`
				Id eventhorizon.UUID `json:"id" eh:"optional"`
			}
		*/
		application := student.NewSchoolApplication()
		application.Id = eventhorizon.NewUUID()
		application.ChurchContactPerson, _ = shared.PersonNameParse(str(user["church_responsible"]))
		application.ChurchContact = person.NewContact()
		application.ChurchContact.Phone = str(user["church_contact"])
		application.ChurchCommitment = boolGerman(user["church_commitment"])
		application.Group = str(user["scool_year"])
	}

	/*

		"scool_year":         "Klasse",
			"biblikum_1":         "Biblikum 1 geschrieben?",
			"biblikum_2":         "Biblikum 2 geschrieben?",
			"paided_last_years":  "Gebühr der Vorjahren bezahlt?",
			"church":             "Gemeinde",
			"church_commitment":  "Gemeinde einverstanden?",
			"church_member":      "Mitglied welcher Gemeinde?",
			"church_services":    "Gemeindedienste",
			"church_responsible": "Pastor / Leitungskreis",
			"church_contact":     "Telefon von Pastor / Leitungskreis",
			"address":            "Adresse",
			"plz":                "PLZ",
			"city":               "Ort",
			"job":                "Beruf",
			"education":          "Bildung",
			"marital_state":      "Familienstand",
			"spirit":             "Geistlicher_Werdegang___

	*/
	for _, profile := range profiles {
		json, _ := json.Marshal(profile)
		jsonStr := string(json)
		println(jsonStr)
		createProfile := profilesUrl + string(profile.Id)
		req, err := http.NewRequest(http.MethodPost, createProfile, bytes.NewBuffer(json))
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		res, getErr := client.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}
		if res.StatusCode != http.StatusOK {
			log.Fatal("Status not ok", res.StatusCode)
		}
	}
}

func timeValue(value interface{}) (ret *time.Time) {
	layout := "2006/02/06"
	t, _ := time.Parse(layout, str(value))
	ret = &t
	return
}

func str(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func boolGerman(value interface{}) bool {
	str := str(value)
	if str == "Ja" {
		return true
	} else {
		return false
	}
}

func UUID(value interface{}) eventhorizon.UUID {
	return eventhorizon.UUID(strings.ToLower(str(value)))
}

func writeInsertInto(users1 []string, users2 []string, users3 []string,
	users4 []string, users5 []string, besucher []string, users map[string]map[string]interface{}) {
	f, _ := os.Create(fmt.Sprintf("%v/users.sql", work))
	defer f.Close()
	w := bufio.NewWriter(f)
	writeInsertIntoForGroup(users1, users, "2017 Klasse 1", w)
	writeInsertIntoForGroup(users2, users, "2016 1. Klasse", w)
	writeInsertIntoForGroup(users3, users, "2016 2. Klasse", w)
	writeInsertIntoForGroup(users4, users, "2016 3. Klasse", w)
	writeInsertIntoForGroup(users5, users, "2016 4. Klasse", w)
	w.Flush()
}

func writeInsertIntoForGroup(userKeys []string, users map[string]map[string]interface{}, group string, w *bufio.Writer) {
	//write data
	w.WriteString("insert into wp_users_bs (user_email, gr) values ")
	for _, userKey := range userKeys {
		user := users[userKey]
		//w.WriteString(fmt.Sprintf(`<%v %v>%v,`, user["first_name"], user["last_name"], user["user_email"]))
		w.WriteString(fmt.Sprintf("('%v', '%v'),", user["user_email"], group))
	}
	w.WriteString(";\n")
}

func writeUpdateRole(users1 []string, users2 []string, users3 []string,
	users4 []string, users5 []string, besucher []string, users map[string]map[string]interface{}) {
	f, _ := os.Create(fmt.Sprintf("%v/update_role.sql", work))
	defer f.Close()
	w := bufio.NewWriter(f)
	fmt.Fprintf(w, "update wp_usermeta set meta_value='1-klasse-2017' where meta_key='role' and user_id in (%v);\n",
		ids(users1, users))
	fmt.Fprintf(w, "update wp_usermeta set meta_value='2-klasse-2017-b1' where meta_key='role' and user_id in (%v);\n",
		ids(users2, users))
	fmt.Fprintf(w, "update wp_usermeta set meta_value='3-klasse-2071-b1' where meta_key='role' and user_id in (%v);\n",
		ids(users3, users))
	fmt.Fprintf(w, "update wp_usermeta set meta_value='4-klasse-2017-b2' where meta_key='role' and user_id in (%v);\n",
		ids(users4, users))
	fmt.Fprintf(w, "update wp_usermeta set meta_value='zusatzjahr-b2' where meta_key='role' and user_id in (%v);\n",
		ids(users5, users))
	fmt.Fprintf(w, "update wp_usermeta set meta_value='besucher' where meta_key='role' and user_id in (%v);\n",
		ids(besucher, users))
	w.Flush()
}

func ids(userKeys []string, users map[string]map[string]interface{}) string {
	w := bytes.NewBufferString("")
	for i, userKey := range userKeys {
		user := users[userKey]
		if i > 0 {
			w.WriteString(",")
		}
		id := str(user["ID"])
		if len(id) == 0 || id == "<nil>" {
			println(userKey)
		}
		w.WriteString(id)
	}
	return w.String()
}

func writeCsvEmailGroup(users1 []string, users2 []string, users3 []string,
	users4 []string, users5 []string, besucher []string, users map[string]map[string]interface{}) {
	f, _ := os.Create(fmt.Sprintf("%v/emails.txt", work))
	defer f.Close()
	w := bufio.NewWriter(f)
	writeCsvEmailsForGroup(users1, users, "2017 Klasse 1", w)
	writeCsvEmailsForGroup(users2, users, "2017 Klasse 2", w)
	writeCsvEmailsForGroup(users3, users, "2017 Klasse 3", w)
	writeCsvEmailsForGroup(users4, users, "2017 Klasse 4", w)
	writeCsvEmailsForGroup(users5, users, "2017 Zusatzklasse", w)
	writeCsvEmailsForGroup(besucher, users, "2017 Besucher", w)
	w.Flush()
}

func writeCsvEmailsForGroup(userKeys []string, users map[string]map[string]interface{}, group string, w *bufio.Writer) {
	//write data
	w.WriteString(group)
	w.WriteString(": ")
	for _, userKey := range userKeys {
		user := users[userKey]
		//w.WriteString(fmt.Sprintf(`<%v %v>%v,`, user["first_name"], user["last_name"], user["user_email"]))
		w.WriteString(fmt.Sprintf(`%v,`, user["user_email"]))
	}
	w.WriteString("\n\n\n\n")
}

func writeCsvData(users1 []string, users2 []string, users3 []string,
	users4 []string, users5 []string, besucher []string, users map[string]map[string]interface{}) {
	f, _ := os.Create(fmt.Sprintf("%v/user_data.csv", work))
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString("Nachname;Vorname;Strasse;Plz;Ort;Geburtstag;Gender;E-Mail;Klasse\n")
	writeCsvEmailContactsForGroup(users1, users, []string{"last_name", "first_name", "address", "plz", "city", "birth_date", "gender", "user_email", "17 Klasse 1"}, w)
	writeCsvEmailContactsForGroup(users2, users, []string{"last_name", "first_name", "address", "plz", "city", "birth_date", "gender", "user_email", "17 Klasse 2"}, w)
	writeCsvEmailContactsForGroup(users3, users, []string{"last_name", "first_name", "address", "plz", "city", "birth_date", "gender", "user_email", "17 Klasse 3"}, w)
	writeCsvEmailContactsForGroup(users4, users, []string{"last_name", "first_name", "address", "plz", "city", "birth_date", "gender", "user_email", "17 Klasse 4"}, w)
	writeCsvEmailContactsForGroup(users5, users, []string{"last_name", "first_name", "address", "plz", "city", "birth_date", "gender", "user_email", "17 Zusatzklasse"}, w)
	writeCsvEmailContactsForGroup(besucher, users, []string{"last_name", "first_name", "address", "plz", "city", "birth_date", "gender", "user_email", "17 Besucher"}, w)
	w.Flush()
}

func writeCsvEmailContacts(users1 []string, users2 []string, users3 []string,
	users4 []string, users5 []string, besucher []string, users map[string]map[string]interface{}) {
	f, _ := os.Create(fmt.Sprintf("%v/googleContacts.csv", work))
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString("Last Name;First Name;Birthday;Gender;Mobile Phone;Company;E-mail Address;Categories\n")
	writeCsvEmailContactsForGroup(users1, users, []string{"last_name", "first_name", "address", "plz", "city", "birth_date", "gender", "phone_number",
		"church", "user_email", "17 Klasse 1"}, w)
	writeCsvEmailContactsForGroup(users2, users, []string{"last_name", "first_name", "address", "plz", "city", "birth_date", "gender", "phone_number",
		"church", "user_email", "17 Klasse 2"}, w)
	writeCsvEmailContactsForGroup(users3, users, []string{"last_name", "first_name", "address", "plz", "city", "birth_date", "gender", "phone_number",
		"church", "user_email", "17 Klasse 3"}, w)
	writeCsvEmailContactsForGroup(users4, users, []string{"last_name", "first_name", "address", "plz", "city", "birth_date", "gender", "phone_number",
		"church", "user_email", "17 Klasse 4"}, w)
	writeCsvEmailContactsForGroup(users5, users, []string{"last_name", "first_name", "address", "plz", "city", "birth_date", "gender", "phone_number",
		"church", "user_email", "17 Zusatzklasse"}, w)
	writeCsvEmailContactsForGroup(besucher, users, []string{"last_name", "first_name", "address", "plz", "city", "birth_date", "gender", "phone_number",
		"church", "user_email", "17 Besucher"}, w)
	w.Flush()
}

func writeCsvEmailContactsForGroup(userKeys []string, users map[string]map[string]interface{}, columns []string, w *bufio.Writer) {
	//write data
	for _, userKey := range userKeys {
		user := users[userKey]
		for i, k := range columns {
			cell := user[k]
			if cell == nil {
				cell = k
			}
			if i > 0 {
				w.WriteString(";")
			}
			w.WriteString(fmt.Sprintf(`%v`, cell))
		}
		w.WriteString("\n")
	}
}

func writeHtmlReport(users1 []string, users2 []string, users3 []string,
	users4 []string, users5 []string, besucher []string, allKeys []string, users map[string]map[string]interface{}) {
	f, _ := os.Create(fmt.Sprintf("%v/bewerbungen.html", work))
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString(`<html>
		<head>
			<style>
				table {
					font-size:16px;
					font-family: "Trebuchet MS", Arial, Helvetica, sans-serif;
					border-collapse: collapse;
					border-spacing: 0;
					width: 100%;
				}

				td, th {
					border: 1px solid #ddd;
					text-align: left;
					padding: 8px;
				}

				tr:nth-child(even){background-color: #f2f2f2}

				th {
					padding-top: 11px;
					padding-bottom: 11px;
					background-color: #4CAF50;
					color: white;
				}

				.spirit {
					width: 100px;
				}

				.photo {
					height:	200px;
				}
			</style>
		</head>`)

	user1Keys := []string{"last_name", "first_name", "birth_date", "gender", "phone_number",
		"church", "church_commitment", "church_member", "church_services", "church_responsible",
		"church_contact", "user_email",
		"address", "plz", "city",
		"job", "education",
		"marital_state",
		//"photo",
		"spirit"}

	user2_3Keys := []string{"last_name", "first_name", "birth_date", "gender", "phone_number",
		"church", "church_commitment", "church_member", "church_services", "church_responsible",
		"church_contact", "user_email",
		"address", "plz", "city",
		"job", "education",
		"marital_state", "biblikum_1", "paided_last_years",
		//"photo",
		"spirit"}
	user4_5Keys := []string{"last_name", "first_name", "birth_date", "gender", "phone_number",
		"church", "church_commitment", "church_member", "church_services", "church_responsible",
		"church_contact", "user_email",
		"address", "plz", "city",
		"job", "education",
		"marital_state", "biblikum_1", "biblikum_2", "paided_last_years",
		//"photo",
		"spirit"}

	/*
		user1Keys := []string{"last_name", "first_name", "church", "phone_number", "user_email", "photo"}
		user2_3Keys := user1Keys
		user4_5Keys := user1Keys
	*/

	w.WriteString("<body>")
	w.WriteString("<h2>1. Klasse</h2>")
	writeHtmlReportForGroup(users1, users, user1Keys, allKeys, w)

	w.WriteString("<h2>2. Klasse</h2>")
	writeHtmlReportForGroup(users2, users, user2_3Keys, allKeys, w)
	w.WriteString("<h2>3. Klasse</h2>")
	writeHtmlReportForGroup(users3, users, user2_3Keys, allKeys, w)
	w.WriteString("<h2>4. Klasse</h2>")
	writeHtmlReportForGroup(users4, users, user4_5Keys, allKeys, w)
	w.WriteString("<h2>Zusatzjahr</h2>")
	writeHtmlReportForGroup(users5, users, user4_5Keys, allKeys, w)
	w.WriteString("<h2>Besucher</h2>")
	writeHtmlReportForGroup(besucher, users, user4_5Keys, allKeys, w)

	w.WriteString("</body>")
	w.WriteString("</html>")
	w.Flush()
}

func writeHtmlReportForGroup(userKeys []string, users map[string]map[string]interface{}, columns []string,
	allKeys []string, w *bufio.Writer) {

	sort.Strings(userKeys)

	w.WriteString("<table>")
	w.WriteString("<thead>")
	//write header
	w.WriteString("<th>Nr.</th>")

	//TEST
	//columns = allKeys

	for _, column := range columns {
		w.WriteString(fmt.Sprintf("<th class=\"%v\">%v</th>", column, de[column]))
		//w.WriteString(fmt.Sprintf("<th class=\"%v\">%v</th>", column, column))
	}
	w.WriteString("</thead>")

	w.WriteString("\n")

	//write data
	for i, userKey := range userKeys {
		user := users[userKey]
		w.WriteString("<tr>")
		w.WriteString(fmt.Sprintf("<td>%v</td>", i+1))
		for _, k := range columns {
			cell := user[k]
			if cell == nil {
				w.WriteString("<td></td>")
			} else if k == "photo" {
				w.WriteString(fmt.Sprintf("<td><img src='img/%v' class='photo'></td>", cell))
			} else {
				w.WriteString(fmt.Sprintf("<td>%v</td>", cell))
			}

		}
		w.WriteString("</tr>")
		w.WriteString("\n")
	}
	w.WriteString("</table>")
}

func toString(value interface{}) string {
	return fmt.Sprint(value)
}

func prepareString(str string) (ret string) {
	ret = strings.Replace(strings.Replace(str, "\r\n", "<br>", -1), "\n", "<br>", -1)
	return
}

func fill(mainKey string, m php_serialize.PhpValue, user map[string]interface{}, keys map[string]string) {
	if a, ok := m.(php_serialize.PhpArray); ok {
		for kk, vv := range a {
			kv := fmt.Sprintf("%v", kk)
			if subMap, ok := vv.(php_serialize.PhpArray); ok {
				fill(kv, subMap, user, keys)
			} else {
				if kv == "0" {
					kv = mainKey
				}
				if user[kv] == nil {
					user[kv] = prepareString(toString(vv))
				}
			}
		}
	} else {
		println(a)
	}
}
