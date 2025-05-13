# Tubes2_BE_FireBoyWaterGirl

## Description
Dalam aplikasi ini, pohon dari resep akan dipetakan menggunakan algoritma DFS atau BFS untuk menemukan jalur antara elemen dasar dan elemen target. 

1. Depth-First Search (DFS) 

Algoritma DFS akan melakukan traversal berdasarkan kedalaman. DFS memulai pencarian dari elemen dasardan mengikuti satu jalur hingga mencapai leaf node atau elemen target, baru kemudian kembali untuk mencari jalur lain. Berikut adalah proses dari algoritma DFS. 

Mulai dari node awal (elemen dasar). 

Telusuri satu jalur hingga kedalaman maksimal yaitu hingga mencapai elemen dasar atau node leaf. 

Setelah mencapai leaf node, kembali (backtrack) dan coba jalur lain untuk menemukan kombinasi elemen untuk membentuk elemen target. 

DFS tidak menjamin jalur terpendek, tapi akan menemukan salah satu jalur ke target. 

2. Breadth-First Search (BFS) 

Algortima BFS akan melakukan tranversal berdasarkan tingkat (level): BFS akan memulai pencarian dari elemen dasar dan menelusuri setiap level secara berurutan (level by level) hingga mencapai elemen target. Berikut adalah proses dari algoritma BFS. 

Mulai dari node awal (elemen dasar). 

Kunjungi node yang terhubung langsung dengan node awal (child nodes). 

Kunjungi semua anak node pada level pertama sebelum beralih ke level kedua, dan seterusnya. 

BFS akan menemukan jalur terpendek dalam pohon (apabila ada beberapa jalur menuju target). 

Pemetaan Masalah Pada Aplikasi: 

Elemen: Setiap elemen (misalnya Mud, Fire, Brick) akan menjadi node dalam pohon. 

Edges: Hubungan antara elemen yang bisa membentuk elemen lainnya (misalnya Mud + Fire menghasilkan Brick) akan menjadi edges yang menghubungkan dua node. 

Root: Elemen target yang ingin dicari akan menjadi root atau tujuan akhir dari pencarian. 


## Getting Started
1. Web telah dideploy sehingga Anda bisa menjalankannya dengan mengakses link berikut
 [Akses Web Frontend](https://fireboy-watergirl-frontend.delightfulglacier-589d39d2.southeastasia.azurecontainerapps.io/)

## Author
Muh. Rusmin Nurwadin (13523068)
Aryo Bama Wiratama (13523088)
Reza Ahmad Syarif (13523119)
