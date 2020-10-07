-- phpMyAdmin SQL Dump
-- version 4.6.6deb5
-- https://www.phpmyadmin.net/
--
-- Client :  localhost:3306
-- Généré le :  Mer 24 Juin 2020 à 15:21
-- Version du serveur :  5.7.30-0ubuntu0.18.04.1
-- Version de PHP :  7.3.17-1+ubuntu18.04.1+deb.sury.org+1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

USE mailing;

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données :  `mailing`
--

-- --------------------------------------------------------

--
-- Structure de la table `broadcast`
--

CREATE TABLE `broadcast` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `description` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Contenu de la table `broadcast`
--

INSERT INTO `broadcast` (`id`, `name`, `description`) VALUES
(1, 'Anciens Heticiens', 'Les anciens');

-- --------------------------------------------------------

--
-- Structure de la table `broadcast_subscriber`
--

CREATE TABLE `broadcast_subscriber` (
  `broadcast_id` int(11) NOT NULL,
  `subscriber_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Contenu de la table `broadcast_subscriber`
--

INSERT INTO `broadcast_subscriber` (`broadcast_id`, `subscriber_id`) VALUES
(1, 1);

-- --------------------------------------------------------

--
-- Structure de la table `client`
--

CREATE TABLE `client` (
  `id` int(11) NOT NULL,
  `mail` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `token` varchar(255) NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Contenu de la table `client`
--

INSERT INTO `client` (`id`, `mail`, `password`, `token`) VALUES
(1, 'wyllismonteiro@gmail.com', '$argon2id$v=19$m=65536,t=3,p=2$0vp4/qRbRTSLSEagVPbuPw$8F14MfBTKOC5ghiWP1N231TrjJAqR8O0UrEaWbeT3vk', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1OTMwMDYxNDMsInVzZXIiOiJXeWxsaXMgTW9udGVpcm8ifQ.oRkp5E98N1DMYkZAFa9hluLTtLNDyIernt3fnPXpnuk');

-- --------------------------------------------------------

--
-- Structure de la table `subscriber`
--

CREATE TABLE `subscriber` (
  `id` int(11) NOT NULL,
  `mail` varchar(100) NOT NULL,
  `name` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Contenu de la table `subscriber`
--

INSERT INTO `subscriber` (`id`, `mail`, `name`) VALUES
(1, 'kevin@gmail.com', 'Kevin');

-- --------------------------------------------------------

--
-- Structure de la table `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `mail` varchar(100) NOT NULL,
  `client_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Contenu de la table `user`
--

INSERT INTO `user` (`id`, `mail`, `client_id`) VALUES
(1, 'test', 1);

--
-- Index pour les tables exportées
--

--
-- Index pour la table `broadcast`
--
ALTER TABLE `broadcast`
  ADD PRIMARY KEY (`id`);

--
-- Index pour la table `broadcast_subscriber`
--
ALTER TABLE `broadcast_subscriber`
  ADD KEY `broadcast_id` (`broadcast_id`),
  ADD KEY `subscriber_id` (`subscriber_id`);

--
-- Index pour la table `client`
--
ALTER TABLE `client`
  ADD PRIMARY KEY (`id`);

--
-- Index pour la table `subscriber`
--
ALTER TABLE `subscriber`
  ADD PRIMARY KEY (`id`);

--
-- Index pour la table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`),
  ADD KEY `client_id` (`client_id`);

--
-- AUTO_INCREMENT pour les tables exportées
--

--
-- AUTO_INCREMENT pour la table `broadcast`
--
ALTER TABLE `broadcast`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
--
-- AUTO_INCREMENT pour la table `client`
--
ALTER TABLE `client`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
--
-- AUTO_INCREMENT pour la table `subscriber`
--
ALTER TABLE `subscriber`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
--
-- AUTO_INCREMENT pour la table `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
--
-- Contraintes pour les tables exportées
--

--
-- Contraintes pour la table `broadcast_subscriber`
--
ALTER TABLE `broadcast_subscriber`
  ADD CONSTRAINT `broadcast_subscriber_ibfk_1` FOREIGN KEY (`broadcast_id`) REFERENCES `broadcast` (`id`),
  ADD CONSTRAINT `broadcast_subscriber_ibfk_2` FOREIGN KEY (`subscriber_id`) REFERENCES `subscriber` (`id`);

--
-- Contraintes pour la table `user`
--
ALTER TABLE `user`
  ADD CONSTRAINT `user_ibfk_1` FOREIGN KEY (`client_id`) REFERENCES `client` (`id`);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
