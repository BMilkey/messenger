package database

import (
	"fmt"

	"github.com/BMilkey/messenger/hlp"
	pgx "github.com/jackc/pgx/v5/pgxpool"
)

func createDB(pool *pgx.Pool, cfg hlp.DatabaseConfig) error {
	querry := fmt.Sprintf(
		`
		CREATE DATABASE %s
		WITH
		OWNER = %s
		ENCODING = 'UTF8'
		LC_COLLATE = 'Russian_Russia.1251'
		LC_CTYPE = 'Russian_Russia.1251'
		LOCALE_PROVIDER = 'libc'
		TABLESPACE = pg_default
		CONNECTION LIMIT = -1
		IS_TEMPLATE = False;
		`,
		cfg.DbName,
		cfg.User)
	return wrapRawExecQuerry(pool, querry)
}

func createTables(pool *pgx.Pool, cfg hlp.DatabaseConfig) error {
	tables := []func(*pgx.Pool, hlp.DatabaseConfig) error{
		createFiles,
		createImages,
		createUsers,
		createChats,
		createChatParticipants,
		createMessages,
		createUserAvatars,
		createImageCalls,
		createFileCalls,
	}

	for _, table := range tables {
		if err := table(pool, cfg); err != nil {
			return err
		}
	}

	return nil
}

func createFiles(pool *pgx.Pool, cfg hlp.DatabaseConfig) error {
	querry := fmt.Sprintf(
		`
		CREATE TABLE IF NOT EXISTS public.files
		(
		    id character varying(128) COLLATE pg_catalog."default" NOT NULL,
		    name character varying(128) COLLATE pg_catalog."default" NOT NULL,
		    path character varying(256) COLLATE pg_catalog."default" NOT NULL,
		    size bigint NOT NULL,
		    CONSTRAINT files_pkey PRIMARY KEY (id),
		    CONSTRAINT unique_file_path UNIQUE (path)
		)

		TABLESPACE pg_default;

		ALTER TABLE IF EXISTS public.files
		    OWNER to %s;
		`,
		cfg.User)
	return wrapRawExecQuerry(pool, querry)
}

func createImages(pool *pgx.Pool, cfg hlp.DatabaseConfig) error {
	querry := fmt.Sprintf(
		`
		CREATE TABLE IF NOT EXISTS public.images
		(
			-- Inherited from table public.files: id character varying(128) COLLATE pg_catalog."default" NOT NULL,
			-- Inherited from table public.files: name character varying(128) COLLATE pg_catalog."default" NOT NULL,
			-- Inherited from table public.files: path character varying(256) COLLATE pg_catalog."default" NOT NULL,
			-- Inherited from table public.files: size bigint NOT NULL,
			CONSTRAINT images_pkey PRIMARY KEY (id),
			CONSTRAINT unique_image_path UNIQUE (path)
		)
			INHERITS (public.files)
		
		TABLESPACE pg_default;
		
		ALTER TABLE IF EXISTS public.images
			OWNER to %s;
		`,
		cfg.User)
	return wrapRawExecQuerry(pool, querry)
}

func createUsers(pool *pgx.Pool, cfg hlp.DatabaseConfig) error {
	querry := fmt.Sprintf(
		`
		CREATE TABLE IF NOT EXISTS public.users
		(
			id character varying(128) COLLATE pg_catalog."default" NOT NULL,
			name character varying(64) COLLATE pg_catalog."default" NOT NULL,
			link character varying(64) COLLATE pg_catalog."default" NOT NULL,
			about character varying(512) COLLATE pg_catalog."default" NOT NULL,
			last_connection timestamp with time zone NOT NULL,
			image_id character varying(128) COLLATE pg_catalog."default",
			CONSTRAINT users_pkey PRIMARY KEY (id),
			CONSTRAINT unique_links UNIQUE (link),
			CONSTRAINT "FK_user_picture_id" FOREIGN KEY (image_id)
				REFERENCES public.images (id) MATCH SIMPLE
				ON UPDATE NO ACTION
				ON DELETE NO ACTION
				NOT VALID
		)
		
		TABLESPACE pg_default;
		
		ALTER TABLE IF EXISTS public.users
			OWNER to %s;
		`,
		cfg.User)
	return wrapRawExecQuerry(pool, querry)
}

func createChats(pool *pgx.Pool, cfg hlp.DatabaseConfig) error {
	querry := fmt.Sprintf(
		`
		CREATE TABLE IF NOT EXISTS public.chats
		(
			id character varying(128) COLLATE pg_catalog."default" NOT NULL,
			link character varying(128) COLLATE pg_catalog."default" NOT NULL,
			title character varying(128) COLLATE pg_catalog."default" NOT NULL,
			created_by_user_id character varying(128) COLLATE pg_catalog."default" NOT NULL,
			create_time timestamp with time zone NOT NULL,
			about character varying(512) COLLATE pg_catalog."default" NOT NULL,
			image_id character varying(128) COLLATE pg_catalog."default",
			CONSTRAINT chats_pkey PRIMARY KEY (id),
			CONSTRAINT unique_chat_link UNIQUE (link),
			CONSTRAINT "FK_chat_image_id" FOREIGN KEY (image_id)
				REFERENCES public.images (id) MATCH SIMPLE
				ON UPDATE NO ACTION
				ON DELETE NO ACTION
				NOT VALID,
			CONSTRAINT "FK_created_by_user_id" FOREIGN KEY (created_by_user_id)
				REFERENCES public.users (id) MATCH SIMPLE
				ON UPDATE NO ACTION
				ON DELETE NO ACTION
				NOT VALID
		)
		
		TABLESPACE pg_default;
		
		ALTER TABLE IF EXISTS public.chats
			OWNER to %s;
		`,
		cfg.User)
	return wrapRawExecQuerry(pool, querry)
}

func createChatParticipants(pool *pgx.Pool, cfg hlp.DatabaseConfig) error {
	querry := fmt.Sprintf(
		`
		CREATE TABLE IF NOT EXISTS public.chat_participants
		(
			chat_id character varying(128) COLLATE pg_catalog."default" NOT NULL,
			user_id character varying(128) COLLATE pg_catalog."default" NOT NULL,
			CONSTRAINT chat_participants_pkey PRIMARY KEY (chat_id, user_id),
			CONSTRAINT "FK_chat_participants_chat_id" FOREIGN KEY (chat_id)
				REFERENCES public.chats (id) MATCH SIMPLE
				ON UPDATE NO ACTION
				ON DELETE NO ACTION,
			CONSTRAINT "FK_chat_participants_user_id" FOREIGN KEY (user_id)
				REFERENCES public.users (id) MATCH SIMPLE
				ON UPDATE NO ACTION
				ON DELETE NO ACTION
		)
		
		TABLESPACE pg_default;
		
		ALTER TABLE IF EXISTS public.chat_participants
			OWNER to %s;
		`,
		cfg.User)
	return wrapRawExecQuerry(pool, querry)
}

func createMessages(pool *pgx.Pool, cfg hlp.DatabaseConfig) error {
	querry := fmt.Sprintf(
		`

		CREATE TABLE IF NOT EXISTS public.messages
		(
			id character varying(128) COLLATE pg_catalog."default" NOT NULL,
			chat_id character varying(128) COLLATE pg_catalog."default" NOT NULL,
			user_id character varying(128) COLLATE pg_catalog."default" NOT NULL,
			create_time timestamp with time zone NOT NULL,
			text character varying(4096) COLLATE pg_catalog."default" NOT NULL,
			reply_message_id character varying(128) COLLATE pg_catalog."default" NOT NULL,
			CONSTRAINT messages_pkey PRIMARY KEY (id),
			CONSTRAINT "FK_messages_chat_id" FOREIGN KEY (chat_id)
				REFERENCES public.chats (id) MATCH SIMPLE
				ON UPDATE NO ACTION
				ON DELETE NO ACTION,
			CONSTRAINT "FK_messages_reply_msg_id" FOREIGN KEY (reply_message_id)
				REFERENCES public.messages (id) MATCH SIMPLE
				ON UPDATE NO ACTION
				ON DELETE NO ACTION
				NOT VALID,
			CONSTRAINT "FK_messages_user_id" FOREIGN KEY (user_id)
				REFERENCES public.users (id) MATCH SIMPLE
				ON UPDATE NO ACTION
				ON DELETE NO ACTION
		)
		
		TABLESPACE pg_default;
		
		ALTER TABLE IF EXISTS public.messages
			OWNER to %s;
		`,
		cfg.User)
	return wrapRawExecQuerry(pool, querry)
}

func createUserAvatars(pool *pgx.Pool, cfg hlp.DatabaseConfig) error {
	querry := fmt.Sprintf(
		`
		CREATE TABLE IF NOT EXISTS public.user_avatars
		(
			user_id character varying(128) COLLATE pg_catalog."default" NOT NULL,
			img_id character varying(128) COLLATE pg_catalog."default" NOT NULL,
			CONSTRAINT user_avatars_pkey PRIMARY KEY (user_id, img_id),
			CONSTRAINT "FK_user_avatars_img_id" FOREIGN KEY (img_id)
				REFERENCES public.images (id) MATCH SIMPLE
				ON UPDATE NO ACTION
				ON DELETE NO ACTION,
			CONSTRAINT "FK_user_avatars_user_id" FOREIGN KEY (user_id)
				REFERENCES public.users (id) MATCH SIMPLE
				ON UPDATE NO ACTION
				ON DELETE NO ACTION
		)
		
		TABLESPACE pg_default;
		
		ALTER TABLE IF EXISTS public.user_avatars
			OWNER to %s;
		`,
		cfg.User)
	return wrapRawExecQuerry(pool, querry)
}

func createImageCalls(pool *pgx.Pool, cfg hlp.DatabaseConfig) error {
	querry := fmt.Sprintf(
		`
		CREATE TABLE IF NOT EXISTS public.image_calls
		(
			message_id character varying(128) COLLATE pg_catalog."default" NOT NULL,
			image_id character varying(128) COLLATE pg_catalog."default" NOT NULL,
			CONSTRAINT image_calls_pkey PRIMARY KEY (message_id, image_id),
			CONSTRAINT "FK_image_calls_img_id" FOREIGN KEY (image_id)
				REFERENCES public.images (id) MATCH SIMPLE
				ON UPDATE NO ACTION
				ON DELETE NO ACTION,
			CONSTRAINT "FK_image_calls_msg_id" FOREIGN KEY (message_id)
				REFERENCES public.messages (id) MATCH SIMPLE
				ON UPDATE NO ACTION
				ON DELETE NO ACTION
		)
		
		TABLESPACE pg_default;
		
		ALTER TABLE IF EXISTS public.image_calls
			OWNER to %s;
		`,
		cfg.User)
	return wrapRawExecQuerry(pool, querry)
}

func createFileCalls(pool *pgx.Pool, cfg hlp.DatabaseConfig) error {
	querry := fmt.Sprintf(
		`
		CREATE TABLE IF NOT EXISTS public.file_calls
		(
			message_id character varying(128) COLLATE pg_catalog."default" NOT NULL,
			file_id character varying(128) COLLATE pg_catalog."default" NOT NULL,
			CONSTRAINT file_calls_pkey PRIMARY KEY (message_id, file_id),
			CONSTRAINT "FK_file_calls_file_id" FOREIGN KEY (file_id)
				REFERENCES public.files (id) MATCH SIMPLE
				ON UPDATE NO ACTION
				ON DELETE NO ACTION,
			CONSTRAINT "FK_file_calls_msg_id" FOREIGN KEY (message_id)
				REFERENCES public.messages (id) MATCH SIMPLE
				ON UPDATE NO ACTION
				ON DELETE NO ACTION
		)
		
		TABLESPACE pg_default;
		
		ALTER TABLE IF EXISTS public.file_calls
			OWNER to %s;
		`,
		cfg.User)
	return wrapRawExecQuerry(pool, querry)
}
