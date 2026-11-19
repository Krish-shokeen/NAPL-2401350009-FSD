(ns ques1.core
  (:gen-class))

(defn -main [& args]

  ;; 1. Java Static Methods
  (println "===== JAVA STATIC METHODS =====")

  (println "Absolute value:" (Math/abs -25))
  (println "Square root:" (Math/sqrt 144))
  (println "Current Time:" (System/currentTimeMillis))


  ;; 2. Creating Java Objects
  (println "\n===== JAVA OBJECTS =====")

  (def arr-list (java.util.ArrayList.))
  (.add arr-list "Apple")
  (.add arr-list "Banana")
  (.add arr-list "Mango")

  (println "ArrayList:" arr-list)
  (println "ArrayList Size:" (.size arr-list))


  ;; 3. Calling Instance Methods
  (println "\n===== INSTANCE METHODS =====")

  (.add arr-list "Orange")
  (println "After adding Orange:" arr-list)

  (println "Element at index 1:" (.get arr-list 1))

  (.remove arr-list 0)
  (println "After removing first element:" arr-list)


  ;; HashMap
  (def java-map (java.util.HashMap.))
  (.put java-map "Name" "Vipul")
  (.put java-map "Roll No" "2401350017")
  (.put java-map "Course" "CSE")

  (println "Java HashMap:" java-map)
  (println "Student Name:" (.get java-map "Name"))


  ;; 4. Clojure and Java Collection Conversion
  (println "\n===== COLLECTION CONVERSION =====")

  (def clojure-vector [10 20 30 40 50])

  (println "Clojure Vector:" clojure-vector)

  ;; Clojure to Java
  (def java-list (java.util.ArrayList. clojure-vector))
  (println "Java ArrayList:" java-list)

  ;; Java to Clojure
  (def converted-vector (vec java-list))
  (println "Converted back to Clojure Vector:" converted-vector)

  (def converted-set (set java-list))
  (println "Converted to Clojure Set:" converted-set))