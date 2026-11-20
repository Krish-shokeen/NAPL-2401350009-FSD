(defproject datapipeline "0.1.0-SNAPSHOT"
  :description "Clojure Data Pipeline with Ring HTTP and SQLite"
  :dependencies [[org.clojure/clojure "1.11.1"]
                 [ring/ring-core "1.12.2"]
                 [ring/ring-jetty-adapter "1.12.2"]
                 [org.clojure/data.json "2.5.1"]
                 [org.xerial/sqlite-jdbc "3.46.1.3"]]
  :main datapipeline.core
  :target-path "target/%s"
  :profiles {:uberjar {:aot :all}})