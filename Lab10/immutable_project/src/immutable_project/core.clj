(ns immutable-project.core
  (:gen-class))

(defn -main [& args]

  ;; 1. LIST
  (println "----- LIST -----")
  (def my-list (list 10 20 30 40))
  (println "Original List:" my-list)

  (def new-list (conj my-list 5))
  (println "After conj:" new-list)

  (println "First element:" (first my-list))
  (println "Count:" (count my-list))
  (println "Original List after operations:" my-list)


  ;; 2. VECTOR
  (println "\n----- VECTOR -----")
  (def my-vector [10 20 30 40])
  (println "Original Vector:" my-vector)

  (def new-vector (conj my-vector 50))
  (println "After conj:" new-vector)

  (println "Get index 2:" (get my-vector 2))
  (println "Count:" (count my-vector))
  (println "Original Vector after operations:" my-vector)


  ;; 3. MAP
  (println "\n----- MAP -----")
  (def my-map {:name "Krish" :age 20 :course "CSE"})
  (println "Original Map:" my-map)

  (def updated-map (assoc my-map :age 21))
  (println "After assoc:" updated-map)

  (def removed-map (dissoc my-map :course))
  (println "After dissoc:" removed-map)

  (println "Get name:" (get my-map :name))
  (println "Original Map after operations:" my-map)


  ;; 4. SET
  (println "\n----- SET -----")
  (def my-set #{10 20 30 40})
  (println "Original Set:" my-set)

  (def new-set (conj my-set 50))
  (println "After conj:" new-set)

  (def removed-set (disj my-set 20))
  (println "After disj:" removed-set)

  (println "Contains 30:" (contains? my-set 30))
  (println "Original Set after operations:" my-set))