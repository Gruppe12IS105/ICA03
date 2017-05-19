# IS-105 ICA03

## ASCII (0x00-0x7F)
Felles: Alle funksjonene kjøres fra hver sin main-fil utenfor pakken.
### Oppgave 1a)
**ICA03 1/1a/** IterateOverASCIIStringLiteral( ) tar en for-løkke over en array av bytes som inneholder alle tegn fra ASCII-tabellen. Deretter skriver den ut i tre forskjellige formater: Hexadesimal, karakteren verdien representerer og binær. Utskriften er godt eksempel på hvordan pakken fmt kan brukes og Printf-funksjonen.
### Oppgave 1b)
**ICA03 1/1b/** GreetingASCII( ) gjør det samme som IterateOverASCIIStringLiteral( ), men her blir utskriften “Hello :-)”.  
### Oppgave 1c)
**ICA03 1/1c/** Eneste nye i denne oppgaven er ascii_test.go som sjekker om alle byte-verdiene i Hello :-) er innenfor ASCII-tabellen. Denne testen vil få PASS.
## Extended ASCII (0x80-0xFF)
Felles: Alle funksjonene kjøres fra hver sin main-fil utenfor pakken.
### Oppgave 2a)
**ICA03 2/2a/** Gjør det samme som koden i 1a) bare fra Extended ASCII. Verdiene skrives ut som hexadesimal, karakteren verdien representerer og binær.
### Oppgave 2b)
**ICA03 2/2b/** GreetingExtendedASCII( ) tar en for-løkke over en array av bytes. Utskriften blir: "Salut, ça va °+) €50" (med anførselstegn).
### Oppgave 2c)
**ICA03 2/2c/** Eneste nye i denne oppgaven er iso_test.go som sjekker om alle verdier i "Salut, ça va °+) €50" er innenfor Extended ASCII. Testen vil få FAIL siden kun € er innenfor.
## Oppgave 3
### Oppgave 3a)
**ICA03-3/3a/** skriver ut en []byte i string-format. Utskriften blir: "½?=? ⌘"
### Oppgave 3b)
**ICA03-3/3b/** Åpner en av tre filer (filename som ikke er kommentert ut) og henter ut informasjon om innholdet og konverterer det til en array av bytes som den returnerer. Derfra skrives file_analysis.go ut byteslicen på hexadesimal.
### Oppgave 3c)
**ICA03-3/3c/** Bytter ut tre verdier fra ASCII-tabellen med Æ, Ø og Å vha Replace( ) og skriver ut resultatet som en string.
## Unicode
### Oppgave 4a)
**ICA03 4/4a/** Tar inn et argument (jp eller is) og skriver ut "nord og sør" på japansk eller islandsk avhengig av hvilket argument som ble skrevet inn. Dette avgjøres av en if-setning.
### Oppgave 4b)
**ICA03 4/4b/** Setter opp en server som som skriver ut tiden og en tekst med html-syntaks for forandring av font. 
