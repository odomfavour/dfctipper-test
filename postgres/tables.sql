drop table if exists deposit;
drop table if exists withdrawal;
drop table if exists reward;
drop table if exists promotion;
drop table if exists account;

create table if not exists account (
    id VARCHAR(256) NOT NULL primary key,
    referral_id VARCHAR(256) NOT NULL,
    balance INT8 NOT NULL,
    username VARCHAR(256) NOT NULL,
    first_name VARCHAR(256) NOT NULL,
    last_name VARCHAR(256) NOT NULL,
    twitter_id int8 NOT NULL,
    telegram_id int8 NOT NULL unique,
    join_at INT8 NOT NULL,
    wallet_address VARCHAR(256) NOT NULL,
    deposit_address VARCHAR(256) NOT NULL,
	current_step INT NOT NULL
);

create table if not exists promotion (
    id serial NOT NULL primary key,
    creator_id VARCHAR(256) NOT NULL REFERENCES account(id),
    created_at INT8 NOT NULL,
    tweet_link VARCHAR(256) NOT NULL,
    reward_count INT NOT NULL,
    reward_per_retweet INT NOT NULL,
    retweet_count INT NOT NULL
);

create table if not exists reward (
    id serial not null primary key,
    user_id VARCHAR(256) not null REFERENCES account(id),
    promotion_id int not null REFERENCES promotion(id),
    date int8 not null,
    amount int8 not null
);

create table if not exists withdrawal (
    id serial not null primary key,
    user_id VARCHAR(256) not null REFERENCES account(id),
    amount int8 not null,
    tx_hash VARCHAR(256) NOT NULL,
    date int8 not null
);

create table if not exists deposit (
    id serial not null primary key,
    amount INT8 not null,
    tx_hash VARCHAR(256) not null,
    user_id VARCHAR(256) not null REFERENCES account(id),
    date int8 not null
);

alter table account add column active int not null default 1;
alter table account add downlines int not null default 0;
alter table account add contest_downline int not null default 0;
