(ns ques2.core
  (:gen-class))

(defn -main [& args]

  ;; Collection 1: Vector
  (println "===== VECTOR =====")

  (def numbers [1 2 3 4 5 6 7 8 9 10])

  (println "Original Vector:" numbers)

  ;; map
  (println "Squares:" (map #(* % %) numbers))

  ;; filter
  (println "Even Numbers:" (filter even? numbers))

  ;; reduce
  (println "Sum:" (reduce + numbers))

  ;; take and drop
  (println "First 5 Numbers:" (take 5 numbers))
  (println "After Dropping 5:" (drop 5 numbers))

  ;; Threading Macro
  (println "Pipeline Result:"
           (->> numbers
                (filter even?)
                (map #(* % %))
                (take 3)))


  ;; Collection 2: List
  (println "\n===== LIST =====")

  (def my-list '(10 15 20 25 30 35 40))

  (println "Original List:" my-list)

  ;; map
  (println "Add 5:" (map #(+ % 5) my-list))

  ;; filter
  (println "Numbers Greater Than 20:" (filter #(> % 20) my-list))

  ;; reduce
  (println "Maximum Number:" (reduce max my-list))

  ;; take and drop
  (println "First 3 Elements:" (take 3 my-list))
  (println "After Dropping 3:" (drop 3 my-list))

  ;; Threading Macro
  (println "Pipeline Result:"
           (->> my-list
                (filter odd?)
                (map #(* % 2))
                (take 3))))