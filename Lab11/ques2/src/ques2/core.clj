(ns ques2.core
  (:gen-class))

(defn divide-numbers [a b]
  (/ a b))

(defn risky-operation []
  (throw (ex-info "Custom exception occurred!"
                  {:type :custom-error
                   :info "Something went wrong"})))

(defn -main [& args]

  (println "===== ARITHMETIC EXCEPTION =====")
  (try
    (divide-numbers 10 0)
    (catch ArithmeticException e
      (println "Caught ArithmeticException:" (.getMessage e)))
    (finally
      (println "Finally executed for arithmetic")))

  (println "\n===== NULL POINTER EXCEPTION =====")
  (try
    (let [text nil]
      (.length text))
    (catch NullPointerException e
      (println "Caught NullPointerException:" (.getMessage e)))
    (finally
      (println "Finally executed for null")))

  (println "\n===== CUSTOM EXCEPTION =====")
  (try
    (risky-operation)
    (catch clojure.lang.ExceptionInfo e
      (println "Exception Message:" (ex-message e))
      (println "Exception Data:" (ex-data e)))
    (finally
      (println "Finally executed for custom exception")))

  (println "\n===== GRACEFUL RECOVERY =====")
  (try
    (throw (ex-info "Invalid input detected"
                    {:input -5
                     :reason "Negative value not allowed"}))
    (catch clojure.lang.ExceptionInfo e
      (println "Error Message:" (ex-message e))
      (println "Error Data:" (ex-data e))
      (println "Recovered Value: 0")
      (println "Program continues normally"))))