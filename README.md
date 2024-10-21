# Entry-Stamp - Eingangsstempel Generieren

Das Programm generiert ein PNG Bild des Eingangsstempels und speichert es im Order in dem es aufgerufen wurde.
Es wird automatisch das heutige Datum verwendet, mithilfe der `-Datum` Option kann auch ein beliebiges gewählt werden. Zusätzlich kann die Schriftgröße mit `-Größe` angegeben werden.

Die Kopf und Fusszeile des Stempels müssen in der _config.yaml_ Datei angegeben werden. Eine Beispieldatei liegt unter _example_config.yaml_, nach der kompilierung sind diese Werte auch ohne das _config.yaml_ gespeichert.

## Benutzung

Das kompilierte programm mit dem namen _stamp_, kann wie folgt ausgeführt werden.
    
    ./stamp     // Mac OS, Linux
    stamp.exe   // Windows

Es stehen folgende Optionen zur Verfügung:

    -Datum string
            Das verwendete Datum im Format 'YYYY-MM-DD'; Ist automatisch als heutiges Datum konfiguriert. (default "2024-10-21")
    -Größe int
            Die verwendete Schriftgröße. (default 18)

Beispielaufruf mit Optionen:

    stamp.exe -Datum 2024-10-19 -Größe 44
    -> Stempel unter '/Users/<user>/eingangsstempel_19_10_2024.png' gespeichert.


## Kompilieren

Um das Programm zu kompilieren muss `go build` verwendet werden, mit der optionalen `-o` Option, kann der Name der Datei gewählt werden.
Dazu muss Golang installiert sein, und der gesamte Projektordner installiert sein. 

    go build -o stamp.exe   // Windows
    go build -o stamp       // Unix


## Kontakt

Carl Steinbach - steinbachcf01@gmail.com

---

Version: _0.1_


