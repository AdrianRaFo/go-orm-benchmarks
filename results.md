# Results

- orm-benchmark (with no flags)

```
   
Reports:

   200 times - Insert
       ent:     0.03s       171805 ns/op    4004 B/op     95 allocs/op
       pgx:     0.03s       172603 ns/op     306 B/op      9 allocs/op
  pgx_pool:     0.04s       181107 ns/op    1140 B/op     14 allocs/op
      sqlc:     0.04s       184210 ns/op    3266 B/op     65 allocs/op x
        pg:     0.04s       187466 ns/op    6357 B/op     11 allocs/op x
      gorp:     0.04s       190069 ns/op    2565 B/op     45 allocs/op x
 gorm_prep:     0.04s       195340 ns/op    6258 B/op     72 allocs/op
       bun:     0.04s       195801 ns/op    5400 B/op     15 allocs/op
     beego:     0.04s       198572 ns/op    2393 B/op     58 allocs/op
      zorm:     0.04s       198639 ns/op    3587 B/op     64 allocs/op
       raw:     0.04s       199231 ns/op    1470 B/op     16 allocs/op
    reform:     0.04s       201333 ns/op    2507 B/op     54 allocs/op
      gorm:     0.04s       207189 ns/op    7824 B/op    107 allocs/op
      sqlx:     0.04s       208335 ns/op     864 B/op     20 allocs/op
 sqlboiler:     0.04s       212149 ns/op    2379 B/op     37 allocs/op
      xorm:     0.04s       212853 ns/op    3472 B/op     90 allocs/op
       dbr:     0.05s       228650 ns/op    2825 B/op     67 allocs/op
      godb:     0.05s       240894 ns/op    4700 B/op    117 allocs/op

   200 times - MultiInsert 100 row
       pgx:     0.08s       407544 ns/op   47820 B/op     33 allocs/op
  pgx_pool:     0.08s       414048 ns/op   48568 B/op     34 allocs/op
       raw:     0.11s       565764 ns/op  212016 B/op    545 allocs/op
     beego:     0.13s       628366 ns/op  178286 B/op   2747 allocs/op
    reform:     0.14s       683034 ns/op  488292 B/op   2356 allocs/op
 gorm_prep:     0.14s       712877 ns/op  254438 B/op   1890 allocs/op
       ent:     0.16s       796276 ns/op  412693 B/op   4209 allocs/op
        pg:     0.17s       873492 ns/op    8976 B/op    113 allocs/op
      sqlx:     0.18s       882789 ns/op  172533 B/op   1553 allocs/op
      xorm:     0.18s       904749 ns/op  248664 B/op   5415 allocs/op
      zorm:     0.18s       912677 ns/op  199704 B/op   2768 allocs/op
       bun:     0.19s       925809 ns/op   42855 B/op    218 allocs/op
      godb:     0.21s      1060524 ns/op  254566 B/op   5893 allocs/op
      gorm:     0.23s      1134404 ns/op  294346 B/op   5231 allocs/op
       dbr:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert

   200 times - Update
       ent:     0.00s        22766 ns/op    4505 B/op     96 allocs/op
       raw:     0.01s        37109 ns/op     650 B/op     11 allocs/op
      sqlx:     0.01s        56891 ns/op     880 B/op     21 allocs/op
       pgx:     0.03s       170695 ns/op     298 B/op      9 allocs/op
      gorp:     0.03s       171328 ns/op    1108 B/op     30 allocs/op
    reform:     0.03s       174682 ns/op    1676 B/op     49 allocs/op
      sqlc:     0.04s       175605 ns/op     778 B/op     12 allocs/op
  pgx_pool:     0.04s       177115 ns/op     317 B/op      9 allocs/op
       bun:     0.04s       183764 ns/op    4736 B/op      6 allocs/op
 gorm_prep:     0.04s       187344 ns/op    5019 B/op     57 allocs/op
       dbr:     0.04s       189574 ns/op    2659 B/op     58 allocs/op
 sqlboiler:     0.04s       191155 ns/op     815 B/op     15 allocs/op
        pg:     0.04s       191673 ns/op     776 B/op     10 allocs/op
      gorm:     0.04s       192180 ns/op    6761 B/op    100 allocs/op
     beego:     0.04s       196893 ns/op    1763 B/op     48 allocs/op
      xorm:     0.04s       206584 ns/op    3953 B/op    133 allocs/op
      zorm:     0.05s       226097 ns/op    2880 B/op     52 allocs/op
      godb:     0.05s       238704 ns/op    5137 B/op    155 allocs/op

   200 times - Read
       pgx:     0.00s        20140 ns/op     800 B/op     19 allocs/op
  pgx_pool:     0.00s        20706 ns/op     993 B/op     20 allocs/op
      sqlc:     0.00s        21748 ns/op    1778 B/op     54 allocs/op
     beego:     0.00s        22169 ns/op    2112 B/op     77 allocs/op
       raw:     0.00s        23521 ns/op    1686 B/op     52 allocs/op
      gorp:     0.00s        24835 ns/op    4070 B/op    204 allocs/op
 gorm_prep:     0.01s        25935 ns/op    4489 B/op     88 allocs/op
 sqlboiler:     0.01s        26839 ns/op    1175 B/op     24 allocs/op
       bun:     0.01s        29542 ns/op    5530 B/op     22 allocs/op
        pg:     0.01s        30710 ns/op     881 B/op     21 allocs/op
       ent:     0.01s        31473 ns/op    5039 B/op    152 allocs/op
    reform:     0.01s        32430 ns/op    2803 B/op     88 allocs/op
       dbr:     0.01s        33310 ns/op    2192 B/op     38 allocs/op
      gorm:     0.01s        33596 ns/op    4855 B/op     98 allocs/op
      zorm:     0.01s        35711 ns/op    3024 B/op     62 allocs/op
      sqlx:     0.01s        46267 ns/op    1984 B/op     44 allocs/op
      xorm:     0.01s        50917 ns/op    4657 B/op    128 allocs/op
      godb:     0.01s        55733 ns/op    4088 B/op    103 allocs/op

   200 times - MultiRead limit 100
    reform:     0.00s        18617 ns/op    2514 B/op     76 allocs/op
       pgx:     0.02s        89032 ns/op   31185 B/op    615 allocs/op
  pgx_pool:     0.02s        95804 ns/op   31255 B/op    615 allocs/op
       raw:     0.02s       105490 ns/op   26291 B/op   1141 allocs/op
        pg:     0.02s       111977 ns/op   27919 B/op    630 allocs/op
      sqlc:     0.03s       126158 ns/op   60359 B/op   1252 allocs/op
      sqlx:     0.03s       135158 ns/op   37488 B/op   1226 allocs/op
      gorp:     0.03s       147342 ns/op   45202 B/op   1503 allocs/op
       dbr:     0.03s       153732 ns/op   30824 B/op   1255 allocs/op
       bun:     0.03s       161925 ns/op   34132 B/op   1125 allocs/op
       ent:     0.03s       165375 ns/op   62498 B/op   2041 allocs/op
 sqlboiler:     0.04s       175757 ns/op   54196 B/op   2260 allocs/op
     beego:     0.04s       180014 ns/op   55315 B/op   3079 allocs/op
      gorm:     0.04s       208261 ns/op   43982 B/op   2092 allocs/op
 gorm_prep:     0.04s       208993 ns/op   43639 B/op   2082 allocs/op
      godb:     0.04s       211833 ns/op   75384 B/op   3085 allocs/op
      zorm:     0.05s       245451 ns/op  161627 B/op   2947 allocs/op
      xorm:     0.06s       278192 ns/op  119582 B/op   4402 allocs/op
```

- orm-benchmark -multi=5

```
  
Reports:

  1000 times - Insert
  pgx_pool:     0.17s       166121 ns/op     471 B/op     10 allocs/op
     beego:     0.17s       171065 ns/op    2396 B/op     58 allocs/op
       pgx:     0.17s       171342 ns/op     291 B/op      9 allocs/op
       raw:     0.17s       172646 ns/op     773 B/op     12 allocs/op
 sqlboiler:     0.17s       173924 ns/op    1652 B/op     33 allocs/op
        pg:     0.18s       177388 ns/op    1913 B/op     11 allocs/op
      sqlc:     0.18s       178946 ns/op    2562 B/op     61 allocs/op
    reform:     0.18s       180915 ns/op    1842 B/op     50 allocs/op
 gorm_prep:     0.18s       182544 ns/op    5410 B/op     67 allocs/op
       bun:     0.18s       182952 ns/op    5102 B/op     15 allocs/op
      gorm:     0.18s       184386 ns/op    7331 B/op    107 allocs/op
       ent:     0.18s       184957 ns/op    4015 B/op     96 allocs/op
       dbr:     0.19s       186136 ns/op    2722 B/op     66 allocs/op
      gorp:     0.19s       191238 ns/op    1864 B/op     41 allocs/op
      sqlx:     0.19s       192098 ns/op     864 B/op     20 allocs/op
      xorm:     0.20s       203199 ns/op    3360 B/op     90 allocs/op
      zorm:     0.21s       212371 ns/op    3577 B/op     65 allocs/op
      godb:     0.23s       225154 ns/op    4560 B/op    116 allocs/op

  1000 times - MultiInsert 100 row
  pgx_pool:     0.40s       399446 ns/op   48145 B/op     34 allocs/op
       pgx:     0.41s       405092 ns/op   47672 B/op     33 allocs/op
       raw:     0.57s       567751 ns/op  211041 B/op    539 allocs/op
     beego:     0.59s       592195 ns/op  178175 B/op   2746 allocs/op
    reform:     0.66s       660097 ns/op  489955 B/op   2356 allocs/op
       ent:     0.81s       805385 ns/op  412346 B/op   4207 allocs/op
      sqlx:     0.85s       850421 ns/op  172151 B/op   1552 allocs/op
        pg:     0.86s       859584 ns/op    4461 B/op    113 allocs/op
       bun:     0.92s       922385 ns/op   42763 B/op    219 allocs/op
      zorm:     0.94s       936723 ns/op  199704 B/op   2768 allocs/op
      xorm:     0.95s       948428 ns/op  248725 B/op   5416 allocs/op
      godb:     1.05s      1048315 ns/op  254455 B/op   5895 allocs/op
 gorm_prep:     1.13s      1127061 ns/op  253945 B/op   1891 allocs/op
      gorm:     1.45s      1449799 ns/op  294851 B/op   5232 allocs/op
      sqlc:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert

  1000 times - Update
       ent:     0.03s        25117 ns/op    4512 B/op     96 allocs/op
       raw:     0.03s        31493 ns/op     648 B/op     11 allocs/op
      sqlx:     0.04s        42845 ns/op     880 B/op     21 allocs/op
  pgx_pool:     0.17s       166637 ns/op     310 B/op      9 allocs/op
     beego:     0.17s       169020 ns/op    1761 B/op     48 allocs/op
       pgx:     0.17s       170831 ns/op     296 B/op      9 allocs/op
 sqlboiler:     0.17s       171517 ns/op     803 B/op     15 allocs/op
      sqlc:     0.17s       172355 ns/op     776 B/op     12 allocs/op
    reform:     0.18s       176540 ns/op    1672 B/op     49 allocs/op
       dbr:     0.18s       181343 ns/op    2659 B/op     58 allocs/op
      gorp:     0.19s       185355 ns/op    1105 B/op     30 allocs/op
       bun:     0.19s       191078 ns/op    4742 B/op      6 allocs/op
        pg:     0.20s       201071 ns/op     776 B/op     10 allocs/op
      zorm:     0.20s       204716 ns/op    2880 B/op     52 allocs/op
      xorm:     0.21s       210443 ns/op    3956 B/op    133 allocs/op
      gorm:     0.21s       211748 ns/op    6773 B/op    100 allocs/op
      godb:     0.22s       216591 ns/op    5144 B/op    155 allocs/op
 gorm_prep:     0.33s       329888 ns/op    5022 B/op     57 allocs/op

  1000 times - Read
  pgx_pool:     0.02s        19402 ns/op     974 B/op     20 allocs/op
       pgx:     0.02s        19644 ns/op     787 B/op     19 allocs/op
      sqlc:     0.02s        21455 ns/op    1776 B/op     54 allocs/op
     beego:     0.02s        23485 ns/op    2111 B/op     77 allocs/op
       raw:     0.02s        23789 ns/op    1668 B/op     52 allocs/op
        pg:     0.02s        24986 ns/op     880 B/op     21 allocs/op
       bun:     0.03s        25435 ns/op    5556 B/op     22 allocs/op
 sqlboiler:     0.03s        26143 ns/op     855 B/op     24 allocs/op
       dbr:     0.03s        26768 ns/op    2192 B/op     38 allocs/op
      gorp:     0.03s        27180 ns/op    4063 B/op    204 allocs/op
      zorm:     0.03s        30188 ns/op    3024 B/op     62 allocs/op
    reform:     0.03s        30394 ns/op    2798 B/op     88 allocs/op
       ent:     0.03s        31593 ns/op    5042 B/op    152 allocs/op
      gorm:     0.03s        34663 ns/op    4835 B/op     98 allocs/op
      sqlx:     0.04s        42627 ns/op    1984 B/op     44 allocs/op
 gorm_prep:     0.04s        44896 ns/op    4446 B/op     88 allocs/op
      xorm:     0.06s        56955 ns/op    4661 B/op    128 allocs/op
      godb:     0.06s        64061 ns/op    4092 B/op    103 allocs/op

  1000 times - MultiRead limit 100
    reform:     0.03s        26987 ns/op    2521 B/op     76 allocs/op
  pgx_pool:     0.09s        90089 ns/op   31255 B/op    615 allocs/op
       pgx:     0.09s        90230 ns/op   31191 B/op    615 allocs/op
       raw:     0.11s       107731 ns/op   26287 B/op   1141 allocs/op
        pg:     0.12s       120323 ns/op   24485 B/op    630 allocs/op
      sqlc:     0.13s       130563 ns/op   60358 B/op   1252 allocs/op
      sqlx:     0.14s       136546 ns/op   37488 B/op   1226 allocs/op
       bun:     0.15s       150600 ns/op   34169 B/op   1125 allocs/op
       dbr:     0.16s       159237 ns/op   30824 B/op   1255 allocs/op
       ent:     0.16s       161048 ns/op   62516 B/op   2041 allocs/op
      gorp:     0.17s       165588 ns/op   45216 B/op   1503 allocs/op
     beego:     0.17s       172228 ns/op   55308 B/op   3079 allocs/op
 sqlboiler:     0.19s       191382 ns/op   54610 B/op   2260 allocs/op
      gorm:     0.22s       222865 ns/op   43956 B/op   2092 allocs/op
      godb:     0.24s       237168 ns/op   75363 B/op   3085 allocs/op
      zorm:     0.27s       268377 ns/op  161625 B/op   2947 allocs/op
 gorm_prep:     0.30s       300904 ns/op   43475 B/op   2082 allocs/op
      xorm:     0.33s       331284 ns/op  119573 B/op   4402 allocs/op
```

- orm-benchmark -multi=10

```
  
Reports:

  2000 times - Insert
       pgx:     0.33s       163781 ns/op     289 B/op      9 allocs/op
  pgx_pool:     0.33s       164437 ns/op     391 B/op      9 allocs/op
       raw:     0.33s       166198 ns/op     696 B/op     11 allocs/op
      gorp:     0.34s       170775 ns/op    1783 B/op     41 allocs/op
 sqlboiler:     0.34s       170838 ns/op    1562 B/op     33 allocs/op
       ent:     0.34s       171143 ns/op    4013 B/op     96 allocs/op
 gorm_prep:     0.35s       173108 ns/op    5297 B/op     67 allocs/op
      sqlc:     0.35s       175168 ns/op    2477 B/op     61 allocs/op
     beego:     0.35s       176820 ns/op    2396 B/op     58 allocs/op
        pg:     0.35s       176911 ns/op     833 B/op     11 allocs/op
       bun:     0.36s       179794 ns/op    5060 B/op     15 allocs/op
    reform:     0.37s       187197 ns/op    1764 B/op     49 allocs/op
       dbr:     0.38s       188386 ns/op    2709 B/op     66 allocs/op
      gorm:     0.38s       190800 ns/op    7278 B/op    107 allocs/op
      sqlx:     0.39s       196700 ns/op     864 B/op     20 allocs/op
      xorm:     0.40s       199717 ns/op    3346 B/op     90 allocs/op
      godb:     0.40s       200391 ns/op    4551 B/op    116 allocs/op
      zorm:     0.42s       211923 ns/op    3576 B/op     65 allocs/op

  2000 times - MultiInsert 100 row
  pgx_pool:     0.80s       399441 ns/op   47759 B/op     33 allocs/op
       pgx:     0.80s       402340 ns/op   47993 B/op     34 allocs/op
       raw:     1.14s       572001 ns/op  211072 B/op    539 allocs/op
    reform:     1.31s       656169 ns/op  488823 B/op   2356 allocs/op
 gorm_prep:     1.39s       695846 ns/op  253865 B/op   1891 allocs/op
       ent:     1.58s       787901 ns/op  411965 B/op   4207 allocs/op
      sqlx:     1.71s       856738 ns/op  171445 B/op   1552 allocs/op
        pg:     1.73s       867073 ns/op    4420 B/op    113 allocs/op
       bun:     1.76s       881596 ns/op   42745 B/op    220 allocs/op
      zorm:     1.80s       900242 ns/op  199704 B/op   2768 allocs/op
      xorm:     1.87s       937249 ns/op  248672 B/op   5415 allocs/op
      gorm:     1.98s       989355 ns/op  293924 B/op   5232 allocs/op
      godb:     2.06s      1031529 ns/op  254569 B/op   5895 allocs/op
     beego:     2.20s      1098266 ns/op  178253 B/op   2747 allocs/op
      gorp:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert

  2000 times - Update
       raw:     0.04s        19663 ns/op     648 B/op     11 allocs/op
       ent:     0.05s        25407 ns/op    4513 B/op     96 allocs/op
      sqlx:     0.09s        43319 ns/op     880 B/op     21 allocs/op
  pgx_pool:     0.34s       168791 ns/op     312 B/op      9 allocs/op
     beego:     0.34s       169908 ns/op    1762 B/op     48 allocs/op
    reform:     0.34s       172032 ns/op    1677 B/op     49 allocs/op
       dbr:     0.35s       175367 ns/op    2659 B/op     58 allocs/op
 gorm_prep:     0.35s       175464 ns/op    5023 B/op     57 allocs/op
      sqlc:     0.36s       178259 ns/op     776 B/op     12 allocs/op
 sqlboiler:     0.36s       179374 ns/op     801 B/op     15 allocs/op
        pg:     0.37s       183881 ns/op     776 B/op     10 allocs/op
      gorp:     0.37s       185232 ns/op    1106 B/op     30 allocs/op
       pgx:     0.37s       186492 ns/op     296 B/op      9 allocs/op
      gorm:     0.38s       190576 ns/op    6771 B/op    100 allocs/op
       bun:     0.39s       196415 ns/op    4745 B/op      6 allocs/op
      xorm:     0.40s       198009 ns/op    3958 B/op    133 allocs/op
      zorm:     0.42s       207892 ns/op    2880 B/op     52 allocs/op
      godb:     0.44s       219409 ns/op    5147 B/op    155 allocs/op

  2000 times - Read
  pgx_pool:     0.04s        18976 ns/op     973 B/op     20 allocs/op
       pgx:     0.04s        19630 ns/op     785 B/op     19 allocs/op
      sqlc:     0.04s        22065 ns/op    1778 B/op     54 allocs/op
     beego:     0.05s        23136 ns/op    2110 B/op     77 allocs/op
       raw:     0.05s        23328 ns/op    1668 B/op     52 allocs/op
        pg:     0.05s        23352 ns/op     880 B/op     21 allocs/op
    reform:     0.05s        24597 ns/op    2795 B/op     88 allocs/op
      gorp:     0.05s        24916 ns/op    4064 B/op    204 allocs/op
 sqlboiler:     0.05s        25622 ns/op     816 B/op     24 allocs/op
 gorm_prep:     0.05s        25818 ns/op    4449 B/op     88 allocs/op
       bun:     0.05s        26034 ns/op    5552 B/op     22 allocs/op
       dbr:     0.06s        27734 ns/op    2192 B/op     38 allocs/op
       ent:     0.07s        32741 ns/op    5042 B/op    152 allocs/op
      gorm:     0.07s        34601 ns/op    4827 B/op     98 allocs/op
      zorm:     0.08s        41525 ns/op    3024 B/op     62 allocs/op
      sqlx:     0.09s        43901 ns/op    1984 B/op     44 allocs/op
      godb:     0.10s        49342 ns/op    4097 B/op    103 allocs/op
      xorm:     0.10s        50495 ns/op    4662 B/op    128 allocs/op

  2000 times - MultiRead limit 100
    reform:     0.04s        20297 ns/op    2521 B/op     76 allocs/op
       pgx:     0.18s        87640 ns/op   31194 B/op    615 allocs/op
  pgx_pool:     0.18s        88807 ns/op   31255 B/op    615 allocs/op
       raw:     0.22s       111961 ns/op   26302 B/op   1141 allocs/op
        pg:     0.24s       119821 ns/op   23965 B/op    630 allocs/op
      sqlc:     0.26s       129827 ns/op   60354 B/op   1252 allocs/op
      sqlx:     0.27s       133546 ns/op   37488 B/op   1226 allocs/op
      gorp:     0.29s       146334 ns/op   45207 B/op   1503 allocs/op
       dbr:     0.32s       159211 ns/op   30824 B/op   1255 allocs/op
       ent:     0.32s       162076 ns/op   62521 B/op   2041 allocs/op
       bun:     0.33s       164331 ns/op   34159 B/op   1125 allocs/op
     beego:     0.37s       182586 ns/op   55312 B/op   3079 allocs/op
 sqlboiler:     0.38s       190952 ns/op   54600 B/op   2260 allocs/op
 gorm_prep:     0.39s       194650 ns/op   43479 B/op   2082 allocs/op
      gorm:     0.40s       201317 ns/op   43874 B/op   2092 allocs/op
      godb:     0.49s       243552 ns/op   75395 B/op   3085 allocs/op
      zorm:     0.55s       272623 ns/op  161625 B/op   2947 allocs/op
      xorm:     0.63s       315018 ns/op  119576 B/op   4402 allocs/op
```

- orm-benchmark -multi=20

```
  
Reports:

  4000 times - Insert
  pgx_pool:     0.65s       163394 ns/op     347 B/op      9 allocs/op
       raw:     0.67s       166635 ns/op     650 B/op     11 allocs/op
       pgx:     0.67s       167891 ns/op     288 B/op      9 allocs/op
     beego:     0.67s       168386 ns/op    2395 B/op     58 allocs/op
       ent:     0.71s       176713 ns/op    4015 B/op     96 allocs/op
 sqlboiler:     0.71s       176838 ns/op    1518 B/op     33 allocs/op
      gorp:     0.71s       177675 ns/op    1742 B/op     41 allocs/op
      sqlc:     0.72s       180140 ns/op    2431 B/op     61 allocs/op
       dbr:     0.72s       180223 ns/op    2702 B/op     66 allocs/op
 gorm_prep:     0.72s       180345 ns/op    5246 B/op     67 allocs/op
        pg:     0.72s       180607 ns/op    1363 B/op     11 allocs/op
      gorm:     0.73s       182236 ns/op    7223 B/op    106 allocs/op
       bun:     0.73s       183610 ns/op    5042 B/op     15 allocs/op
    reform:     0.74s       183920 ns/op    1722 B/op     49 allocs/op
      sqlx:     0.76s       190075 ns/op     864 B/op     20 allocs/op
      zorm:     0.78s       195280 ns/op    3576 B/op     65 allocs/op
      godb:     0.80s       200131 ns/op    4544 B/op    116 allocs/op
      xorm:     0.82s       203957 ns/op    3341 B/op     90 allocs/op

  4000 times - MultiInsert 100 row
       pgx:     1.59s       397191 ns/op   47916 B/op     33 allocs/op
  pgx_pool:     1.63s       408019 ns/op   47604 B/op     33 allocs/op
       raw:     2.33s       582743 ns/op  210415 B/op    538 allocs/op
     beego:     2.40s       600450 ns/op  178027 B/op   2746 allocs/op
    reform:     2.70s       674997 ns/op  490375 B/op   2357 allocs/op
 gorm_prep:     2.75s       686621 ns/op  253992 B/op   1891 allocs/op
       ent:     3.08s       770484 ns/op  412414 B/op   4206 allocs/op
      sqlx:     3.36s       840553 ns/op  171194 B/op   1552 allocs/op
        pg:     3.41s       852126 ns/op    4137 B/op    113 allocs/op
       bun:     3.58s       894447 ns/op   42794 B/op    220 allocs/op
      xorm:     3.75s       938739 ns/op  248863 B/op   5416 allocs/op
      zorm:     3.78s       945608 ns/op  199704 B/op   2768 allocs/op
      gorm:     4.09s      1021945 ns/op  294349 B/op   5232 allocs/op
      godb:     4.18s      1043923 ns/op  254586 B/op   5895 allocs/op
       dbr:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert

  4000 times - Update
       raw:     0.08s        19847 ns/op     649 B/op     11 allocs/op
       ent:     0.12s        30405 ns/op    4513 B/op     96 allocs/op
      sqlx:     0.17s        43549 ns/op     880 B/op     21 allocs/op
       pgx:     0.67s       168586 ns/op     296 B/op      9 allocs/op
 gorm_prep:     0.70s       174733 ns/op    5025 B/op     57 allocs/op
  pgx_pool:     0.70s       175874 ns/op     311 B/op      9 allocs/op
    reform:     0.71s       177797 ns/op    1679 B/op     49 allocs/op
 sqlboiler:     0.71s       178067 ns/op     803 B/op     15 allocs/op
     beego:     0.71s       178378 ns/op    1762 B/op     48 allocs/op
      sqlc:     0.72s       178858 ns/op     777 B/op     12 allocs/op
       dbr:     0.72s       180344 ns/op    2659 B/op     58 allocs/op
      gorp:     0.73s       181614 ns/op    1106 B/op     30 allocs/op
       bun:     0.76s       189905 ns/op    4746 B/op      6 allocs/op
        pg:     0.78s       196129 ns/op    1058 B/op     10 allocs/op
      xorm:     0.80s       199248 ns/op    3959 B/op    133 allocs/op
      zorm:     0.82s       205628 ns/op    2880 B/op     52 allocs/op
      gorm:     0.83s       206351 ns/op    6771 B/op    100 allocs/op
      godb:     0.87s       217509 ns/op    5147 B/op    155 allocs/op

  4000 times - Read
       pgx:     0.08s        19142 ns/op     786 B/op     19 allocs/op
  pgx_pool:     0.08s        20215 ns/op     971 B/op     20 allocs/op
       raw:     0.09s        22525 ns/op    1667 B/op     52 allocs/op
        pg:     0.09s        22932 ns/op     882 B/op     21 allocs/op
      sqlc:     0.10s        23763 ns/op    1779 B/op     54 allocs/op
     beego:     0.10s        25482 ns/op    2107 B/op     77 allocs/op
       bun:     0.10s        25663 ns/op    5554 B/op     22 allocs/op
    reform:     0.10s        25763 ns/op    2797 B/op     88 allocs/op
 sqlboiler:     0.11s        26300 ns/op     809 B/op     24 allocs/op
 gorm_prep:     0.11s        26689 ns/op    4450 B/op     88 allocs/op
       dbr:     0.12s        29942 ns/op    2192 B/op     38 allocs/op
      gorp:     0.13s        32126 ns/op    4064 B/op    204 allocs/op
      zorm:     0.13s        33154 ns/op    3024 B/op     62 allocs/op
       ent:     0.13s        33349 ns/op    5042 B/op    152 allocs/op
      gorm:     0.15s        38276 ns/op    4818 B/op     98 allocs/op
      sqlx:     0.18s        43806 ns/op    1984 B/op     44 allocs/op
      godb:     0.21s        51327 ns/op    4107 B/op    103 allocs/op
      xorm:     0.21s        53655 ns/op    4666 B/op    128 allocs/op

  4000 times - MultiRead limit 100
    reform:     0.09s        22510 ns/op    2521 B/op     76 allocs/op
       pgx:     0.37s        91896 ns/op   31204 B/op    615 allocs/op
  pgx_pool:     0.38s        93853 ns/op   31247 B/op    615 allocs/op
       raw:     0.43s       108653 ns/op   26290 B/op   1141 allocs/op
        pg:     0.47s       116350 ns/op   23678 B/op    630 allocs/op
      sqlc:     0.54s       134766 ns/op   60368 B/op   1252 allocs/op
      sqlx:     0.56s       140693 ns/op   37488 B/op   1226 allocs/op
      gorp:     0.62s       153797 ns/op   45201 B/op   1503 allocs/op
       dbr:     0.62s       155744 ns/op   30824 B/op   1255 allocs/op
       bun:     0.63s       158244 ns/op   34167 B/op   1125 allocs/op
       ent:     0.64s       161083 ns/op   62509 B/op   2041 allocs/op
     beego:     0.74s       185503 ns/op   55290 B/op   3079 allocs/op
 sqlboiler:     0.77s       192948 ns/op   54111 B/op   2260 allocs/op
 gorm_prep:     0.79s       198592 ns/op   43531 B/op   2082 allocs/op
      gorm:     0.86s       214488 ns/op   43832 B/op   2092 allocs/op
      godb:     1.01s       251434 ns/op   75402 B/op   3085 allocs/op
      zorm:     1.06s       264525 ns/op  161625 B/op   2947 allocs/op
      xorm:     1.29s       321340 ns/op  119618 B/op   4402 allocs/op
```
