/*
 * Copyright 2018 The CovenantSQL Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"database/sql"
	"flag"

	"github.com/CovenantSQL/CovenantSQL/client"
	log "github.com/Sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	var config, password, dsn string

	flag.StringVar(&config, "config", "./conf/config.yaml", "config file path")
	flag.StringVar(&dsn, "dsn", "", "database url")
	flag.StringVar(&password, "password", "", "master key password for covenantsql")
	flag.Parse()

	err := client.Init(config, []byte(password))
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("covenantsql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DROP TABLE IF EXISTS testSimple;")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE testSimple ( indexedColumn, nonIndexedColumn );")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE INDEX testIndexedColumn ON testSimple ( indexedColumn );")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO testSimple VALUES(?, ?)", 4, 400)
	if err != nil {
		log.Fatal(err)
	}

	row := db.QueryRow("SELECT nonIndexedColumn FROM testSimple LIMIT 1")

	var result int
	err = row.Scan(&result)
	if err != nil {
		log.Fatal(err)
	}
	log.Debugf("result %d", result)

	err = db.Close()
}
