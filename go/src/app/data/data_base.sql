-- Adminer 4.6.3 PostgreSQL dump

DROP TABLE IF EXISTS "tab_cambio";
DROP SEQUENCE IF EXISTS tab_cambio_tc_id_seq;
CREATE SEQUENCE tab_cambio_tc_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1;

CREATE TABLE "public"."tab_cambio" (
    "tc_id" integer DEFAULT nextval('tab_cambio_tc_id_seq') NOT NULL,
    "tc_data" date NOT NULL,
    "tc_valor_origem" numeric(14,2) NOT NULL,
    "tc_valor_destino" numeric(14,2) NOT NULL,
    "tc_tipo_cambio" integer NOT NULL,
    CONSTRAINT "tab_cambio_pkey" PRIMARY KEY ("tc_id"),
    CONSTRAINT "tab_cambio_tc_tipo_cambio_fkey" FOREIGN KEY (tc_tipo_cambio) REFERENCES tab_tipo_cambio(ttc_id) NOT DEFERRABLE
) WITH (oids = false);

COMMENT ON COLUMN "public"."tab_cambio"."tc_data" IS 'Dia em que o cambio foi persistido';

COMMENT ON COLUMN "public"."tab_cambio"."tc_valor_origem" IS 'Valor da moeda de origem';

COMMENT ON COLUMN "public"."tab_cambio"."tc_valor_destino" IS 'Valor da moeda de destino';

COMMENT ON COLUMN "public"."tab_cambio"."tc_tipo_cambio" IS 'Indentificador do tipo de cambio';


DROP TABLE IF EXISTS "tab_tipo_cambio";
DROP SEQUENCE IF EXISTS tab_tipo_cambio_ttc_id_seq;
CREATE SEQUENCE tab_tipo_cambio_ttc_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1;

CREATE TABLE "public"."tab_tipo_cambio" (
    "ttc_id" integer DEFAULT nextval('tab_tipo_cambio_ttc_id_seq') NOT NULL,
    "ttc_moeda_origem" text DEFAULT '' NOT NULL,
    "ttc_moeda_destino" text DEFAULT '' NOT NULL,
    CONSTRAINT "constraintname" UNIQUE ("ttc_id"),
    CONSTRAINT "tab_tipo_cambio_pkey" PRIMARY KEY ("ttc_id")
) WITH (oids = false);

COMMENT ON COLUMN "public"."tab_tipo_cambio"."ttc_moeda_origem" IS 'Descricao da origem da moeda';

COMMENT ON COLUMN "public"."tab_tipo_cambio"."ttc_moeda_destino" IS 'Descricao do destino da moeda';

INSERT INTO "tab_tipo_cambio" ("ttc_id", "ttc_moeda_origem", "ttc_moeda_destino") VALUES
(1,	'USA',	'BRL'),
(2,	'EUR',	'BRL');

-- 2018-11-16 18:29:17.563854+00
