# ppcsv

Egy speciális excel **568588.xlsx** fájl tartalmából készít egy adatbázis számára feldolgozható csv fájlt

### Működés leírás

Nincs szükség telepítésre, bárhonnan indítható.
Indítás parancssorból windows 10 környezetben alapértelmezett paraméterekkel: 
Start menü -> cmd

```cmd

c:\path\to\ppcsv.exe

```
Ebben az esetben az **568588.xlsx** fájlnak a programmal azonos könyvtárban kell lennie
Két darab csv fájl lesz az eredmény:
- **failedlist.csv** : ide kerülnek a további feldolgozást igénylő rekordok
- **readylist.csv** : feldolgozásra kész rekordok
Mezőlista:
```
Name       string //Név és vezetéknév
Pcode      string //Irányítószám - parse
City       string //Helység - parse
Address    string //Cím - parse
DefAdr     string //Eredeti cím
Email      string //E-mail
Tel        string //Telefonszáma
Device     string //Készülék típusa
Sn         string //S / N készülék
Srvname    string //Szerviznév
Buyingdate string //Vásárlás dátuma
Startdate  string //Készülék kezdési dátuma
```
A csv fájlok karakterkódolása: utf8 
a mezőelválasztó karakter: ';'
nincs szöveg megjelölés.

Egyelőre a cím mezőben levő szövegek vannak vizsgálva és szétválasztva és
a failed listába azok a rekordok kerülnek, amelyeknél a helység mező üres, vagyis nem sikerült az eredeti szövegből kinyerni a helység nevet.
Mind a két fájlban szerepel az eredeti cím mező is, hogy könnyebb legyen beazonosítani a tévesztéseket. 
Az adatbázis betöltés előtt törölni kell az eredeti cím mezőt.

#### Parancssori kapcsolók

A -h kapcsoló a konzolra írja a további lehetőségeket:

```cmd

c:\path\to\ppcsv.exe -h

Usage of C:\path\to\ppcsv.exe
  -source
        az xlsx forrás fájl helye (default ".\568588.xlsx")
```