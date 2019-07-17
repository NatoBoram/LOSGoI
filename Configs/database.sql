use `LOSGoI`;
-- Builds
drop table if exists `builds`;
create table if not exists `builds`(
  `device` varchar(16) not null,
  `date` date not null,
  `datetime` datetime not null,
  `filename` varchar(64) not null,
  `filepath` varchar(128) not null,
  `sha1` varchar(40) not null,
  `sha256` varchar(64) not null,
  `size` INTEGER not null,
  `type` varchar(16) not null,
  `version` varchar(8) not null,
  `ipfs` varchar(128) not null primary key
);
-- Max
create
or replace view `builds_max` as
select
  `device`,
  max(`datetime`) as `datetime`
from
  `builds`
group by
  `device`;
-- Latest Builds
  create
  or replace view `builds_latest` as
select
  `builds`.`device`,
  `builds`.`date`,
  `builds`.`datetime`,
  `builds`.`filename`,
  `builds`.`filepath`,
  `builds`.`sha1`,
  `builds`.`sha256`,
  `builds`.`size`,
  `builds`.`type`,
  `builds`.`version`,
  `builds`.`ipfs`
from
  `builds`
  join `builds_max`
where
  `builds`.`device` = `builds_max`.`device`
  and `builds`.`datetime` = `builds_max`.`datetime`;
-- Remove garbage from builds
update
  `builds`
set
  `ipfs` = replace(ipfs, char(13), '');
update
  `builds`
set
  `ipfs` = replace(ipfs, char(10), '');