
(ns datapipeline.core
  (:require [ring.adapter.jetty :as jetty]
            [clojure.data.json :as json])
  (:gen-class))

(def db-url "jdbc:sqlite:datapipeline.db")

(def dataset
  [{:id 1 :name "Krish" :marks 92}
   {:id 2 :name "Vipul" :marks 85}
   {:id 3 :name "Abhay" :marks 65}
   {:id 4 :name "Vansh" :marks 78}
   {:id 5 :name "Tarun" :marks 55}])

;; Sequence functions: map, filter and reduce
(defn process-data [data]
  (let [mapped (map #(assoc % :grade
                            (if (>= (:marks %) 80)
                              "A"
                              "B"))
                    data)

        filtered (filter #(>= (:marks %) 70) mapped)

        result (reduce
                (fn [acc student]
                  (conj acc student))
                []
                filtered)]
    result))

;; Create database table
(defn create-table []
  (try
    (Class/forName "org.sqlite.JDBC")
    (with-open [conn (java.sql.DriverManager/getConnection db-url)
                stmt (.createStatement conn)]
      (.execute stmt
                "CREATE TABLE IF NOT EXISTS students
         (id INTEGER PRIMARY KEY,
          name TEXT,
          marks INTEGER,
          grade TEXT)"))
    (println "Database table created successfully")
    (catch Exception e
      (println "Database error:" (.getMessage e)))))

;; Write processed data to database
(defn write-data [data]
  (try
    (with-open [conn (java.sql.DriverManager/getConnection db-url)
                stmt (.prepareStatement conn
                                        "INSERT OR REPLACE INTO students
                  (id, name, marks, grade)
                  VALUES (?, ?, ?, ?)")]
      (doseq [student data]
        (.setInt stmt 1 (:id student))
        (.setString stmt 2 (:name student))
        (.setInt stmt 3 (:marks student))
        (.setString stmt 4 (:grade student))
        (.executeUpdate stmt)))
    (println "Data written successfully")
    (catch Exception e
      (println "Write error:" (.getMessage e)))))

;; Read data from database
(defn read-data []
  (try
    (with-open [conn (java.sql.DriverManager/getConnection db-url)
                stmt (.createStatement conn)
                rs (.executeQuery stmt "SELECT * FROM students")]
      (loop [result []]
        (if (.next rs)
          (recur
           (conj result
                 {:id (.getInt rs "id")
                  :name (.getString rs "name")
                  :marks (.getInt rs "marks")
                  :grade (.getString rs "grade")}))
          result)))
    (catch Exception e
      (println "Read error:" (.getMessage e))
      [])))

;; Ring HTTP handler
(defn handler [request]
  (case (:uri request)

    "/"
    {:status 200
     :headers {"Content-Type" "text/plain"}
     :body "Clojure Data Pipeline is running!"}

    "/api/data"
    {:status 200
     :headers {"Content-Type" "application/json"}
     :body (json/write-str (read-data))}

    "/api/summary"
    (let [data (read-data)
          total (count data)
          average (if (pos? total)
                    (/ (reduce + (map :marks data)) total)
                    0)]
      {:status 200
       :headers {"Content-Type" "application/json"}
       :body (json/write-str
              {:total-students total
               :average-marks average
               :students data})})

    {:status 404
     :headers {"Content-Type" "text/plain"}
     :body "Route not found"}))

;; Main function
(defn -main [& args]
  (create-table)
  (let [processed (process-data dataset)]
    (write-data processed))
  (println "Starting Ring HTTP server on port 3000...")
  (jetty/run-jetty handler {:port 3000 :join? true}))
