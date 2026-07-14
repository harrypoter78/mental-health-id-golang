-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 14 Jul 2026 pada 05.07
-- Versi server: 10.4.32-MariaDB
-- Versi PHP: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `gangguanmental`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `cache`
--

CREATE TABLE `cache` (
  `key` varchar(255) NOT NULL,
  `value` mediumtext NOT NULL,
  `expiration` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `cache_locks`
--

CREATE TABLE `cache_locks` (
  `key` varchar(255) NOT NULL,
  `owner` varchar(255) NOT NULL,
  `expiration` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `failed_jobs`
--

CREATE TABLE `failed_jobs` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `uuid` varchar(255) NOT NULL,
  `connection` text NOT NULL,
  `queue` text NOT NULL,
  `payload` longtext NOT NULL,
  `exception` longtext NOT NULL,
  `failed_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `gejala`
--

CREATE TABLE `gejala` (
  `id_gejala` int(3) NOT NULL,
  `kode_gejala` varchar(3) NOT NULL,
  `nama_gejala` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data untuk tabel `gejala`
--

INSERT INTO `gejala` (`id_gejala`, `kode_gejala`, `nama_gejala`) VALUES
(1, 'G1', 'Kesulitan tidur'),
(2, 'G2', 'Mendengar suara aneh'),
(3, 'G3', 'Sering atau mudah menangis'),
(4, 'G4', 'Kehilangan minat untuk melakukan aktifitas'),
(5, 'G5', 'Emosi menjadi datar'),
(6, 'G6', 'Ingatan terganggu'),
(7, 'G7', 'Menjauh dari lingkungan sosial'),
(8, 'G8', 'Pikiran dan berbicara kacau'),
(9, 'G9', 'Rasa takut dan khawatir berlebihan'),
(10, 'G10', 'Mimpi buruk'),
(11, 'G11', 'Sering merasa sedih'),
(12, 'G12', 'Mempercayai sesuatu yang tidak nyata'),
(13, 'G13', 'Sulit mengendalikan emosi'),
(14, 'G14', 'Diliputi perasaan bersalah berlebihan'),
(15, 'G15', 'Perasaan bermusuhan'),
(16, 'G16', 'Menghindari sebuah tempat atau objek'),
(17, 'G17', 'Kehilangan motivasi'),
(18, 'G18', 'Sering cemas'),
(19, 'G19', 'Mood swing'),
(20, 'G20', 'Perasaan putus asa'),
(21, 'G21', 'Kurangnya daya ingat'),
(22, 'G22', 'Bicara terlalu cepat'),
(23, 'G23', 'Gangguan pernapasan'),
(24, 'G24', 'Gerakan tubuh dan pikiran yang lambat'),
(25, 'G25', 'Merasa kelebihan berat badan'),
(26, 'G26', 'Makan dalam jumlah besar dan dikeluarkan secara paksa'),
(27, 'G27', 'Memiliki obsesi konstan terhadap sesuatu'),
(28, 'G28', 'Melakukan aksi tertentu secara berulang untuk meredakan kecemasan'),
(29, 'G29', 'Takut kotor atau terkena penyakit'),
(30, 'G30', 'Sangat menginginkan segala sesuatu tersusun selaras atau teratur'),
(31, 'G31', 'Suka atau berkeinginan untuk mengumpulkan barang-barang bekas yang Anda temukan'),
(32, 'G32', 'Memiliki obsesi akan kalori dan lemak yang terkandung pada makanan'),
(33, 'G33', 'Memiliki perubahan periode pada saat sering sekali makan atau bahkan tidak sama sekali');

-- --------------------------------------------------------

--
-- Struktur dari tabel `jobs`
--

CREATE TABLE `jobs` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `queue` varchar(255) NOT NULL,
  `payload` longtext NOT NULL,
  `attempts` tinyint(3) UNSIGNED NOT NULL,
  `reserved_at` int(10) UNSIGNED DEFAULT NULL,
  `available_at` int(10) UNSIGNED NOT NULL,
  `created_at` int(10) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `job_batches`
--

CREATE TABLE `job_batches` (
  `id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `total_jobs` int(11) NOT NULL,
  `pending_jobs` int(11) NOT NULL,
  `failed_jobs` int(11) NOT NULL,
  `failed_job_ids` longtext NOT NULL,
  `options` mediumtext DEFAULT NULL,
  `cancelled_at` int(11) DEFAULT NULL,
  `created_at` int(11) NOT NULL,
  `finished_at` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `migrations`
--

CREATE TABLE `migrations` (
  `id` int(10) UNSIGNED NOT NULL,
  `migration` varchar(255) NOT NULL,
  `batch` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data untuk tabel `migrations`
--

INSERT INTO `migrations` (`id`, `migration`, `batch`) VALUES
(1, '2026_05_13_124051_create_admin_table', 0),
(2, '2026_05_13_124051_create_gejala_table', 0),
(3, '2026_05_13_124051_create_penyakit_table', 0),
(4, '2026_05_13_124051_create_riwayat_table', 0),
(5, '2026_05_13_124051_create_rule_table', 0),
(6, '2026_05_13_124054_add_foreign_keys_to_rule_table', 0),
(7, '2026_05_13_124413_create_gejala_table', 0),
(8, '2026_05_13_124413_create_penyakit_table', 0),
(9, '2026_05_13_124413_create_riwayat_table', 0),
(10, '2026_05_13_124413_create_rule_table', 0),
(11, '2026_05_13_124416_add_foreign_keys_to_rule_table', 0),
(12, '2026_05_13_124513_create_gejala_table', 0),
(13, '2026_05_13_124513_create_penyakit_table', 0),
(14, '2026_05_13_124513_create_riwayat_table', 0),
(15, '2026_05_13_124513_create_rule_table', 0),
(16, '2026_05_13_124516_add_foreign_keys_to_rule_table', 0),
(17, '0001_01_01_000000_create_users_table', 1),
(18, '0001_01_01_000001_create_cache_table', 1),
(19, '0001_01_01_000002_create_jobs_table', 1),
(20, '2026_05_14_000001_add_role_to_users', 2),
(21, '2026_05_14_000002_modify_riwayat_table', 2);

-- --------------------------------------------------------

--
-- Struktur dari tabel `password_reset_tokens`
--

CREATE TABLE `password_reset_tokens` (
  `email` varchar(255) NOT NULL,
  `token` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `penyakit`
--

CREATE TABLE `penyakit` (
  `id_penyakit` int(3) NOT NULL,
  `kode_penyakit` varchar(3) NOT NULL,
  `nama_penyakit` varchar(50) NOT NULL,
  `deskripsi` text DEFAULT NULL,
  `solusi_obat` text DEFAULT NULL,
  `solusi_lain` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data untuk tabel `penyakit`
--

INSERT INTO `penyakit` (`id_penyakit`, `kode_penyakit`, `nama_penyakit`, `deskripsi`, `solusi_obat`, `solusi_lain`) VALUES
(1, 'K1', 'Skizofrenia', 'Skizofrenia merupakan penyakit Kesehatan mental yang menyebabkan pasien mengalami halusinasi, mendengar suara aneh dll. Pasien skizofrenia belum memiliki obat khusus tetapi hanya obat untuk mengurangi efek halusinasi. ', 'Obat-obatan berguna untuk menangani halusinasi dan delusi, dokter akan meresepkan obat antipsikotik seperti, Chlorpromazine, Quetiapine, Aripiprazole, Clozapine', 'Psikoterapi pertama terapi individual psikiater akan mengajarkan keluarga dan teman pasien bagaimana berinteraksi dengan pasien, kedua terapi perilaku kognitif berguna mengubah perilaku dan pola pikir pasien, ketiga terapi remediasi kognitif berguna untuk mengajarkan pasien cara memahami lingkungan sosial, meningkatkan kemampuan pasien dalam memperhatikan atau mengingat sesuatu, dan mengendalikan pola pikirnya, keempat terapi elektrokonvulsif berguna untuk meredakan keinginan bunuh diri, mengatasi gejala depresi berat, dan menangani psikosis.'),
(2, 'K2', 'Post Traumatic Stress Disorder (PTSD)', 'Pengobatan PTSD bertujuan untuk meredakan respons emosi pasien dan mengajarkan pasien cara mengendalikan diri dengan baik ketika teringat pada kejadian traumatis.', 'Obat-obatan yang diberikan tergantung pada gejala yang dialami pasien, adalah Antidepresan untuk mengatasi depresi, seperti sertraline dan paroxetine, Anticemas untuk mengatasi kecemasan, seperti benzodiazepine, alprazolam (Xanax), chlordiazepoxide (Librium), clonazepam (Klonopin), diazepam(Valium), dan lorazepam, dan Prazosin untuk mencegah mimpi buruk.', 'Psikoterapi, pertama terapi perilaku kognitif berguna untuk mengenali dan mengubah pola pikir pasien yang negatif menjadi positif, kedua terapi eksposur berguna untuk membantu pasien menghadapi keadaan dan ingatan yang memicu trauma secara efektif, ketiga Eye movement desensitization and reprocessing (EMDR), yaitu kombinasi terapi eksposur dan teknik gerakan mata untuk mengubah respons pasien saat teringat kejadian traumatis.'),
(3, 'K3', 'Depression', 'Depresi adalah gangguan suasana hati yang menyebabkan seseorang teru merasa sedih dan kehilangan minat. Kondisi ini lebih dari sekadar perasaan sedih yang normalnya dialami orang-orang dengan kondisi mentalnya sehat. Ini karena perasaan sedih sangat sulit untuk disingkirkan sehingga terus menerus menghantui. ', 'Obat-obatan yang digunakan untuk mengobati depresi adalah Obat antidepresan, seperti escitalopram, paroxetine, sertraline, fluoxetine, citalopram, venlafaxine, duloxetine, dan bupropion.', 'Psikoterapi, pertama Cognitive behavior therapy (CBT) membantu pengidap melepaskan pikiran dan perasaan negatif, serta menggantinya dengan respon positif, kedua Problem-solving therapy (PST) untuk meningkatkan kemampuan pengidap menghadapi pengalaman yang memicu rasa tertekan, ketiga Interpersonal therapy (IPT) untuk membantu mengatasi masalah yang muncul saat berhubungan dengan orang lain,  keempat Terapi kejut listrik atau electroconvulsive therapy (ECT) untuk pengidap depresi yang tidak membaik setelah diberi obat-obatan, mengalami gejala psikosis, serta pengidap yang mencoba bunuh diri.'),
(4, 'K4', 'Bipolar Disorder', 'Gangguan bipolar adalah kondisi seseorang yang mengalami perubahan suasana hati secara fluktuatif dan drastis, misalnya tiba-tiba menjadi sangat bahagia dari yang sebelumnya murung. Nama lain dari gangguan bipolar adalah manik depresif.', 'Obat obatan yang berguna untuk menyembuhkan bipolar disorder adalah Moodstabilizer seperti lithium, lamotrigine, dan carbamazepine. Antikonvulsan seperti asam valproat. Antipsikotik seperti aripiprazole, olanzapine, quetiapine, dan risperidone. Antidepresan seperti escitalopram, fluoxetine, dan sertraline', 'Psikoterapi pertama Interpersonal and social rhythm therapy (IPSRT).  terfokus pada kestabilan ritme aktivitas sehari-hari, seperti waktu untuk tidur, bangun, hingga makan. Teraturnya ritme dalam beraktivitas mampu membantu pasien untuk mengendalikan gejala gangguan bipolar. Kedua Cognitive behavioral therapy (CBT).  membantu pasien dalam mendeteksi hal yang dapat memicu munculnya gejala gangguan bipolar, sehingga hal tersebut dapat diganti dengan sesuatu yang positif. Ketiga Psychoeducation. Dokter akan mengedukasi pasien dengan hal-hal yang perlu diketahui terkait kondisi yang tengah diderita. Dengan begitu, pasien dapat dengan sendirinya mengidentifikasi penyebab munculnya gejala, menghindarinya, dan membuat strategi penanganan ketika gejala gangguan bipolar muncul.'),
(5, 'K5', 'Paranoria', 'Gangguan kepribadian paranoid adalah jenis gangguan kepribadian eksentrik di mana pengidapnya memiliki rasa curiga dan tidak percaya yang tak ada hentinya terhadap orang lain.', 'Obat-obatan yang dapat membantu penyembuhan yaitu Antipsikosis atipikal, Antipsikosis konvensional, obat penenang.', 'Melakukan pemeriksaann psikolog atau psoloater dengan cara ngobrol atau wawancara dengan penderita serta menjalankan terapi.'),
(6, 'K6', 'Eating Disorder', 'Gangguan makan adalah gangguan mental saat mengonsumsi makanan. Penderita gangguan ini dapat mengonsumsi terlalu sedikit atau terlalu banyak makanan, dan terobsesi pada berat badan atau bentuk tubuhnya.', 'Obat-obatan yang dapat membantu penyembuhan yaitu antidepresan, antikonvulsan, atau anti ADHD bisa mengurangi gejala binge eating. Lisdexamfetamine dimesylate, obat anti-ADHD, adalah obat pertama yang disetujui FDA untuk mengatasi binge eating sedang sampai berat.', 'Melakukan terapi Cognitive Behavioral Therapy (CBT), Interpersonal Psychotherapy (IPT), atau Terapi Penurunan Berat Badan.'),
(7, 'K7', 'Obsessive Compulsive Disorder (OCD)', 'Obsessive compulsive disorder (OCD) adalah gangguan mental yang menyebabkan penderitanya merasa harus melakukan suatu tindakan secara berulang. Bila tidak dilakukan, penderita OCD akan diliputi kecemasan atau ketakutan.', 'Obat antidepresan diberikan bila terapi perilaku kognitif tidak membantu meredakan gejala, atau bila gejala yang dialami cukup parah. Manfaat antidepresan akan terasa setelah 3 bulan pemakaian. Namun pada banyak kasus, pasien perlu mengonsumsi obat ini sampai setidaknya 1 tahun.\r\nJenis obat antidepresan yang umum digunakan untuk mengatasi OCD antara lain Fluoxetine, Fluvoxamine, dan Sertraline\r\n', 'Melakukan terapi perilaku kognitif, pasien akan dihadapkan pada kondisi yang sering kali dihindarinya. Misalnya, psikiater akan meminta penderita yang takut kuman penyakit untuk menyentuh tanah, kemudian mengajarkan cara mengatasi rasa takutnya tersebut. Terapi perilaku kognitif bisa dilakukan secara individu atau berkelompok.'),
(8, 'K8', 'Anxiety Disorder', 'Rasa cemas atau anxiety adalah hal yang normal dirasakan ketika seseorang menghadapi situasi atau mendengar berita yang menimbulkan rasa takut atau khawatir. Namun, anxiety perlu diwaspadai jika muncul tanpa sebab atau sulit dikendalikan, karena bisa jadi hal tersebut disebabkan oleh gangguan kecemasan.', 'Obat-obatan yang dapat membantu penyembuhan yaitu antidepresan, pregabalin, dan benzodiazepine.', 'Pengobatan untuk gangguan kecemasan umum meliputi 2 langkah, yaitu melalui terapi prilaku kognitif (CBT) dan obat-obatan. Kedua langkah ini biasanya akan dikombinasikan sesuai dengan kebutuhan pasien.');

-- --------------------------------------------------------

--
-- Struktur dari tabel `riwayat`
--

CREATE TABLE `riwayat` (
  `id_riwayat` int(3) NOT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL,
  `nama_penyakit` varchar(50) NOT NULL,
  `tanggal` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data untuk tabel `riwayat`
--

INSERT INTO `riwayat` (`id_riwayat`, `user_id`, `nama_penyakit`, `tanggal`) VALUES
(2, NULL, 'Post Traumatic Stress Disorder (PTSD)', '2021-05-19'),
(3, NULL, 'Post Traumatic Stress Disorder (PTSD)', '2021-05-19'),
(4, NULL, 'Post Traumatic Stress Disorder (PTSD)', '2021-05-19'),
(5, NULL, 'Post Traumatic Stress Disorder (PTSD)', '2021-05-19'),
(6, NULL, 'Post Traumatic Stress Disorder (PTSD)', '2021-05-19'),
(7, NULL, 'Post Traumatic Stress Disorder (PTSD)', '2021-05-19'),
(8, NULL, 'Penyakit Tidak Diketahui', '2021-05-19'),
(9, NULL, 'Penyakit Tidak Diketahui', '2021-05-19'),
(10, NULL, 'Depression', '2021-05-19'),
(12, NULL, 'Skizofrenia', '2021-05-20'),
(13, NULL, 'Skizofrenia', '2021-05-20'),
(14, NULL, 'Skizofrenia', '2021-05-20'),
(15, NULL, 'Depression', '2021-05-20'),
(16, NULL, 'Bipolar Disorder', '2021-05-20'),
(17, NULL, 'Post Traumatic Stress Disorder (PTSD)', '2021-05-29'),
(18, NULL, 'Penyakit Tidak Diketahui', '2021-05-29'),
(19, NULL, 'Penyakit Tidak Diketahui', '2021-05-29'),
(20, NULL, 'Penyakit Tidak Diketahui', '2021-05-29'),
(21, NULL, 'Penyakit Tidak Diketahui', '2021-05-29'),
(22, NULL, 'Penyakit Tidak Diketahui', '2021-05-29'),
(23, NULL, 'Penyakit Tidak Diketahui', '2021-05-29'),
(24, NULL, 'Penyakit Tidak Diketahui', '2021-05-29'),
(25, NULL, 'Penyakit Tidak Diketahui', '2021-05-29'),
(26, NULL, 'Penyakit Tidak Diketahui', '2021-05-29'),
(27, NULL, 'Skizofrenia', '2021-05-30'),
(28, NULL, 'Skizofrenia', '2021-05-30'),
(30, NULL, 'Skizofrenia', '2026-05-13'),
(31, NULL, 'Eating Disorder', '2026-05-14'),
(32, 1, 'Skizofrenia', '2026-05-15'),
(33, 3, 'Eating Disorder', '2026-05-15'),
(34, 3, 'Skizofrenia', '2026-05-17'),
(36, 1, 'Skizofrenia', '2026-05-19'),
(37, 1, 'Skizofrenia', '2026-05-19'),
(38, 1, 'Skizofrenia', '2026-05-19'),
(39, 1, 'Obsessive Compulsive Disorder (OCD)', '2026-07-14');

-- --------------------------------------------------------

--
-- Struktur dari tabel `rule`
--

CREATE TABLE `rule` (
  `id_rule` int(3) NOT NULL,
  `kode_rule` int(3) NOT NULL,
  `kode_penyakit` varchar(3) NOT NULL,
  `kode_gejala` varchar(3) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data untuk tabel `rule`
--

INSERT INTO `rule` (`id_rule`, `kode_rule`, `kode_penyakit`, `kode_gejala`) VALUES
(1, 1, 'K1', 'G1'),
(2, 1, 'K1', 'G2'),
(3, 1, 'K1', 'G5'),
(4, 1, 'K1', 'G7'),
(5, 1, 'K1', 'G8'),
(6, 1, 'K1', 'G12'),
(7, 2, 'K2', 'G1'),
(8, 2, 'K2', 'G4'),
(9, 2, 'K2', 'G6'),
(10, 2, 'K2', 'G9'),
(11, 2, 'K2', 'G10'),
(12, 2, 'K2', 'G13'),
(13, 2, 'K2', 'G16'),
(14, 3, 'K3', 'G1'),
(15, 3, 'K3', 'G4'),
(16, 3, 'K3', 'G11'),
(17, 3, 'K3', 'G14'),
(18, 3, 'K3', 'G20'),
(19, 4, 'K8', 'G1'),
(20, 4, 'K8', 'G9'),
(21, 4, 'K8', 'G13'),
(22, 4, 'K8', 'G18'),
(23, 4, 'K8', 'G23'),
(24, 5, 'K5', 'G2'),
(25, 5, 'K5', 'G7'),
(26, 5, 'K5', 'G9'),
(27, 5, 'K5', 'G12'),
(28, 5, 'K5', 'G15'),
(29, 5, 'K5', 'G18'),
(30, 5, 'K5', 'G23'),
(31, 5, 'K5', 'G24'),
(32, 6, 'K4', 'G3'),
(33, 6, 'K4', 'G12'),
(34, 6, 'K4', 'G14'),
(35, 6, 'K4', 'G17'),
(36, 6, 'K4', 'G19'),
(37, 6, 'K4', 'G21'),
(38, 6, 'K4', 'G22'),
(39, 7, 'K6', 'G9'),
(40, 7, 'K6', 'G13'),
(41, 7, 'K6', 'G18'),
(42, 7, 'K6', 'G25'),
(43, 7, 'K6', 'G26'),
(44, 7, 'K6', 'G32'),
(45, 7, 'K6', 'G33'),
(46, 8, 'K7', 'G18'),
(47, 8, 'K7', 'G27'),
(48, 8, 'K7', 'G28'),
(49, 8, 'K7', 'G29'),
(50, 8, 'K7', 'G30'),
(51, 8, 'K7', 'G31');

-- --------------------------------------------------------

--
-- Struktur dari tabel `sessions`
--

CREATE TABLE `sessions` (
  `id` varchar(255) NOT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL,
  `ip_address` varchar(45) DEFAULT NULL,
  `user_agent` text DEFAULT NULL,
  `payload` longtext NOT NULL,
  `last_activity` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data untuk tabel `sessions`
--

INSERT INTO `sessions` (`id`, `user_id`, `ip_address`, `user_agent`, `payload`, `last_activity`) VALUES
('awNCvfOXDlYuhs93bzJFnARn8rEOppnWambgswLp', 1, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/148.0.0.0 Safari/537.36', 'YTo0OntzOjY6Il90b2tlbiI7czo0MDoiTm1ybXhRTERpMGJScU00SlZmTzljc1pGdnlMc3FaUk54eWxselRLNyI7czo5OiJfcHJldmlvdXMiO2E6Mjp7czozOiJ1cmwiO3M6Mzc6Imh0dHA6Ly8xMjcuMC4wLjE6ODAwMC9hZG1pbi9kYXNoYm9hcmQiO3M6NToicm91dGUiO3M6MTU6ImFkbWluLmRhc2hib2FyZCI7fXM6NjoiX2ZsYXNoIjthOjI6e3M6Mzoib2xkIjthOjA6e31zOjM6Im5ldyI7YTowOnt9fXM6NTA6ImxvZ2luX3dlYl81OWJhMzZhZGRjMmIyZjk0MDE1ODBmMDE0YzdmNThlYTRlMzA5ODlkIjtpOjE7fQ==', 1779024730),
('j0w0kXG41Lq0p0oOhjwHPsFMjECoJzNnVeOfr6sI', 1, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Code/1.120.0 Chrome/142.0.7444.265 Electron/39.8.8 Safari/537.36', 'YTo0OntzOjY6Il90b2tlbiI7czo0MDoidlZWUnRwNlFHZUZLOVBnamNIWVJjeXlpYzlFSFdnczdXclRzdFNHYSI7czo5OiJfcHJldmlvdXMiO2E6Mjp7czozOiJ1cmwiO3M6MzM6Imh0dHA6Ly8xMjcuMC4wLjE6ODAwMC9hZG1pbi91c2VycyI7czo1OiJyb3V0ZSI7czoxNzoiYWRtaW4udXNlcnMuaW5kZXgiO31zOjY6Il9mbGFzaCI7YToyOntzOjM6Im9sZCI7YTowOnt9czozOiJuZXciO2E6MDp7fX1zOjUwOiJsb2dpbl93ZWJfNTliYTM2YWRkYzJiMmY5NDAxNTgwZjAxNGM3ZjU4ZWE0ZTMwOTg5ZCI7aToxO30=', 1779025757),
('N7K5WBfp7ubbLpRKxcufPmCAdYNYwXRXuZc18pDr', 1, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/148.0.0.0 Safari/537.36', 'YTo0OntzOjY6Il90b2tlbiI7czo0MDoiUGtqUzllY2dEWkFQSlhUeVAwNVNDUk80cXVBQU1QTFR6TXV2ZjhYNCI7czo5OiJfcHJldmlvdXMiO2E6Mjp7czozOiJ1cmwiO3M6NDE6Imh0dHA6Ly8xMjcuMC4wLjE6ODAwMC9hZG1pbi9nZWphbGEvY3JlYXRlIjtzOjU6InJvdXRlIjtzOjE5OiJhZG1pbi5nZWphbGEuY3JlYXRlIjt9czo2OiJfZmxhc2giO2E6Mjp7czozOiJvbGQiO2E6MDp7fXM6MzoibmV3IjthOjA6e319czo1MDoibG9naW5fd2ViXzU5YmEzNmFkZGMyYjJmOTQwMTU4MGYwMTRjN2Y1OGVhNGUzMDk4OWQiO2k6MTt9', 1779205753),
('UOCguqsvlQZFJJP92zBOGBxOrQS4kvhLcIZpFZJi', NULL, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Code/1.120.0 Chrome/142.0.7444.265 Electron/39.8.8 Safari/537.36', 'YTozOntzOjY6Il90b2tlbiI7czo0MDoiaENZZEIxbkRrS29TZVNZaWc3UUpheTFHYVRPMGN3MUJXalBpUzdDWiI7czo5OiJfcHJldmlvdXMiO2E6Mjp7czozOiJ1cmwiO3M6MjE6Imh0dHA6Ly8xMjcuMC4wLjE6ODAwMCI7czo1OiJyb3V0ZSI7czoxNToiZGlhZ25vc2lzLmluZGV4Ijt9czo2OiJfZmxhc2giO2E6Mjp7czozOiJvbGQiO2E6MDp7fXM6MzoibmV3IjthOjA6e319fQ==', 1779198649);

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `role` enum('admin','user') NOT NULL DEFAULT 'user',
  `email_verified_at` timestamp NULL DEFAULT NULL,
  `password` varchar(255) NOT NULL,
  `remember_token` varchar(100) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `role`, `email_verified_at`, `password`, `remember_token`, `created_at`, `updated_at`) VALUES
(1, 'Administrator', 'admin@example.com', 'admin', '2026-05-14 02:59:58', '$2y$12$JRppI/oDt6fzl2JHloSg7u.Z8w89vwZPC8QXfEA6aU/QyjsrWwA9.', NULL, '2026-05-14 02:59:58', '2026-05-17 06:48:39'),
(2, 'Budi Admin', 'admin@contoh.com', 'admin', NULL, '$2y$12$PZVl8QcT/XZrgk8f3AB8Eu0p5UcZgdMF18jvUodI06763uDF8jAqi', NULL, '2026-05-14 12:06:06', '2026-05-14 12:06:06'),
(3, 'Siti User', 'siti@contoh.com', 'user', NULL, '$2y$12$PZVl8QcT/XZrgk8f3AB8Eu0p5UcZgdMF18jvUodI06763uDF8jAqi', NULL, '2026-05-14 12:06:06', '2026-05-14 12:06:06'),
(4, 'Agus User', 'agus@contoh.com', 'user', NULL, '$2y$12$PZVl8QcT/XZrgk8f3AB8Eu0p5UcZgdMF18jvUodI06763uDF8jAqi', NULL, '2026-05-14 12:06:06', '2026-05-14 12:06:06'),
(5, 'DummyAcc1', 'dummyacc1@gmail.com', 'user', NULL, '$2y$12$hFFW73Zr0Eal7fForC5cseyuLb3btYV00G9NHnn1mcrOLw6JI3fOG', NULL, '2026-05-14 18:08:40', '2026-05-14 18:08:40'),
(6, 'Tes', 'tes123@gmail.com', 'user', NULL, '$2y$12$HsyNnl2AVhqgLw2c17AxU..F0434H9iucRUisbSAWJM1ub7jxg3pi', NULL, '2026-05-17 04:06:10', '2026-05-17 04:06:10'),
(7, 'Dimas Pratama', 'dimas@example.com', 'user', '2026-05-17 03:00:00', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', NULL, '2026-05-17 03:00:00', '2026-05-17 03:00:00'),
(8, 'Ayu Lestari', 'ayu@example.com', 'user', '2026-05-17 03:05:00', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', NULL, '2026-05-17 03:05:00', '2026-05-17 03:05:00'),
(9, 'Rizky Maulana', 'rizky.m@example.com', 'user', '2026-05-17 03:10:00', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', NULL, '2026-05-17 03:10:00', '2026-05-17 03:10:00'),
(10, 'Fitriani', 'fitriani@example.com', 'user', '2026-05-17 03:15:00', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', NULL, '2026-05-17 03:15:00', '2026-05-17 03:15:00'),
(11, 'Hendra Wijaya', 'hendra.w@example.com', 'user', '2026-05-17 03:20:00', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', NULL, '2026-05-17 03:20:00', '2026-05-17 03:20:00'),
(12, 'Rina Marlina', 'rina.marlina@example.com', 'user', '2026-05-17 03:25:00', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', NULL, '2026-05-17 03:25:00', '2026-05-17 03:25:00'),
(13, 'Dedi Saputra', 'dedi.s@example.com', 'user', '2026-05-17 03:30:00', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', NULL, '2026-05-17 03:30:00', '2026-05-17 03:30:00'),
(14, 'Nisa Ramadhani', 'nisa.r@example.com', 'user', '2026-05-17 03:35:00', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', NULL, '2026-05-17 03:35:00', '2026-05-17 03:35:00'),
(15, 'Eko Susanto', 'eko.susanto@example.com', 'user', '2026-05-17 03:40:00', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', NULL, '2026-05-17 03:40:00', '2026-05-17 03:40:00');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `cache`
--
ALTER TABLE `cache`
  ADD PRIMARY KEY (`key`),
  ADD KEY `cache_expiration_index` (`expiration`);

--
-- Indeks untuk tabel `cache_locks`
--
ALTER TABLE `cache_locks`
  ADD PRIMARY KEY (`key`),
  ADD KEY `cache_locks_expiration_index` (`expiration`);

--
-- Indeks untuk tabel `failed_jobs`
--
ALTER TABLE `failed_jobs`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `failed_jobs_uuid_unique` (`uuid`);

--
-- Indeks untuk tabel `gejala`
--
ALTER TABLE `gejala`
  ADD PRIMARY KEY (`id_gejala`),
  ADD KEY `kode_gejala` (`kode_gejala`);

--
-- Indeks untuk tabel `jobs`
--
ALTER TABLE `jobs`
  ADD PRIMARY KEY (`id`),
  ADD KEY `jobs_queue_index` (`queue`);

--
-- Indeks untuk tabel `job_batches`
--
ALTER TABLE `job_batches`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `migrations`
--
ALTER TABLE `migrations`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `password_reset_tokens`
--
ALTER TABLE `password_reset_tokens`
  ADD PRIMARY KEY (`email`);

--
-- Indeks untuk tabel `penyakit`
--
ALTER TABLE `penyakit`
  ADD PRIMARY KEY (`id_penyakit`),
  ADD KEY `kode_penyakit` (`kode_penyakit`);

--
-- Indeks untuk tabel `riwayat`
--
ALTER TABLE `riwayat`
  ADD PRIMARY KEY (`id_riwayat`),
  ADD KEY `riwayat_user_id_foreign` (`user_id`);

--
-- Indeks untuk tabel `rule`
--
ALTER TABLE `rule`
  ADD PRIMARY KEY (`id_rule`),
  ADD KEY `FK_rule_penyakit` (`kode_penyakit`),
  ADD KEY `FK_rule_gejala` (`kode_gejala`);

--
-- Indeks untuk tabel `sessions`
--
ALTER TABLE `sessions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `sessions_user_id_index` (`user_id`),
  ADD KEY `sessions_last_activity_index` (`last_activity`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `users_email_unique` (`email`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `failed_jobs`
--
ALTER TABLE `failed_jobs`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `gejala`
--
ALTER TABLE `gejala`
  MODIFY `id_gejala` int(3) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=38;

--
-- AUTO_INCREMENT untuk tabel `jobs`
--
ALTER TABLE `jobs`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `migrations`
--
ALTER TABLE `migrations`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=22;

--
-- AUTO_INCREMENT untuk tabel `penyakit`
--
ALTER TABLE `penyakit`
  MODIFY `id_penyakit` int(3) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT untuk tabel `riwayat`
--
ALTER TABLE `riwayat`
  MODIFY `id_riwayat` int(3) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=40;

--
-- AUTO_INCREMENT untuk tabel `rule`
--
ALTER TABLE `rule`
  MODIFY `id_rule` int(3) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=53;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `riwayat`
--
ALTER TABLE `riwayat`
  ADD CONSTRAINT `riwayat_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;

--
-- Ketidakleluasaan untuk tabel `rule`
--
ALTER TABLE `rule`
  ADD CONSTRAINT `FK_rule_gejala` FOREIGN KEY (`kode_gejala`) REFERENCES `gejala` (`kode_gejala`),
  ADD CONSTRAINT `FK_rule_penyakit` FOREIGN KEY (`kode_penyakit`) REFERENCES `penyakit` (`kode_penyakit`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
