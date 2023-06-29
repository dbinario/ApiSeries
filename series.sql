-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Versión del servidor:         8.0.30 - MySQL Community Server - GPL
-- SO del servidor:              Win64
-- HeidiSQL Versión:             12.1.0.6537
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Volcando estructura de base de datos para series
CREATE DATABASE IF NOT EXISTS `series` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `series`;

-- Volcando estructura para procedimiento series.actualizarVistos
DELIMITER //
CREATE PROCEDURE `actualizarVistos`()
BEGIN

UPDATE temporadas A

INNER JOIN 

(SELECT id_serie,numero_temporada,COUNT(*) AS vistos FROM episodios WHERE estado='1'
GROUP BY id_serie,numero_temporada) B

ON A.id_serie=B.id_serie AND A.numero_temporada=B.numero_temporada

SET episodios_vistos=B.vistos;

UPDATE temporadas SET faltan_ver=numero_episodios-episodios_vistos;


UPDATE series A

INNER JOIN 

(SELECT id_serie,SUM(episodios_vistos) AS vistos FROM temporadas
GROUP BY id_serie) B

ON A.id_serie=B.id_serie

SET episodios_vistos=B.vistos;

UPDATE series A

INNER JOIN 

(SELECT id_serie,SUM(numero_episodios) AS TOTAL FROM temporadas WHERE estado_temporada IN (1,2) GROUP BY id_serie) B

ON A.id_serie=B.id_serie

SET episodios_activos=B.TOTAL;


UPDATE series SET faltan_ver=episodios_activos-episodios_vistos;

END//
DELIMITER ;

-- Volcando estructura para procedimiento series.eliminarSerie
DELIMITER //
CREATE PROCEDURE `eliminarSerie`(
	IN `idserie` INT
)
BEGIN

DELETE FROM series WHERE id_serie=idserie;
DELETE FROM temporadas WHERE id_serie=idserie;
DELETE FROM capitulos WHERE id_serie=idserie;

END//
DELIMITER ;

-- Volcando estructura para procedimiento series.eliminarTemporada
DELIMITER //
CREATE PROCEDURE `eliminarTemporada`(
	IN `idserie` INT,
	IN `temporada` INT
)
BEGIN

DELETE FROM temporadas WHERE id_serie=idserie AND numero_temporada=temporada;
DELETE FROM episodios WHERE id_serie=idserie AND numero_temporada=temporada;


UPDATE series A

INNER JOIN 

(SELECT id_serie,SUM(numero_episodios) AS episodios FROM temporadas
GROUP BY id_serie) B

ON A.id_serie=B.id_serie

SET A.numero_episodios=B.episodios,episodios_vistos=0,faltan_ver=0;


CALL actualizarVistos();

END//
DELIMITER ;

-- Volcando estructura para tabla series.episodios
CREATE TABLE IF NOT EXISTS `episodios` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_serie` int DEFAULT NULL,
  `numero_temporada` int DEFAULT NULL,
  `numero_episodio` int DEFAULT NULL,
  `nombre_episodio` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `resumen` text COLLATE utf8mb4_general_ci,
  `estado` tinyint DEFAULT '0',
  `fecha_estado` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1562 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- La exportación de datos fue deseleccionada.

-- Volcando estructura para tabla series.series
CREATE TABLE IF NOT EXISTS `series` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_serie` int NOT NULL DEFAULT '0',
  `nombre_serie` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `numero_temporadas` int NOT NULL DEFAULT '0',
  `numero_episodios` int DEFAULT NULL,
  `episodios_activos` int DEFAULT '0',
  `episodios_vistos` int DEFAULT '0',
  `faltan_ver` int DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- La exportación de datos fue deseleccionada.

-- Volcando estructura para tabla series.temporadas
CREATE TABLE IF NOT EXISTS `temporadas` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_serie` int DEFAULT NULL,
  `numero_temporada` int DEFAULT NULL,
  `nombre_temporada` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `numero_episodios` int DEFAULT NULL,
  `episodios_vistos` int DEFAULT '0',
  `faltan_ver` int DEFAULT '0',
  `ultimo_visto` int DEFAULT '0',
  `estado_temporada` tinyint DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=68 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- La exportación de datos fue deseleccionada.

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
