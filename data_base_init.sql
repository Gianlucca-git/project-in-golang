drop table if exists public.identification_type cascade;
create table if not exists public.identification_type (
                                     id serial primary key ,
                                     abbreviation varchar(5) unique not null ,
                                     name varchar unique not null
);
INSERT INTO public.identification_type (id, abbreviation, name) VALUES (1, 'CC', 'Cedula de Ciudadania');
INSERT INTO public.identification_type (id, abbreviation, name) VALUES (2, 'TI', 'Tarjeta de Identidad');
INSERT INTO public.identification_type (id, abbreviation, name) VALUES (3, 'CE', 'Cedula de Extranjeria');
INSERT INTO public.identification_type (id, abbreviation, name) VALUES (4, 'P', 'Pasaporte');
------------------------------------------------------------------------------------------------------------
drop table if exists public.department cascade;
create table if not exists public.department (
                            id serial primary key ,
                            abbreviation varchar(5) unique not null ,
                            name varchar unique not null
);
INSERT INTO public.department (id, abbreviation, name) VALUES (1, 'ADM', 'Administracion');
INSERT INTO public.department (id, abbreviation, name) VALUES (2, 'FIN', 'Financiera');
INSERT INTO public.department (id, abbreviation, name) VALUES (3, 'C', 'COMPRAS');
INSERT INTO public.department (id, abbreviation, name) VALUES (4, 'INF', 'Infraestructura');
INSERT INTO public.department (id, abbreviation, name) VALUES (5, 'OP', 'Operacion');
INSERT INTO public.department (id, abbreviation, name) VALUES (6, 'TH', 'Talento Humano');
INSERT INTO public.department (id, abbreviation, name) VALUES (7, 'SV', 'Servicios Varios');
------------------------------------------------------------------------------------------------------------
drop table if exists public.countries cascade;
create table if not exists public.countries
(
    id serial primary key,
    abbreviation varchar(5) unique not null ,
    name varchar unique not null,
    domain varchar unique not null
);
INSERT INTO public.countries (id, abbreviation, name, domain) VALUES (1, 'COL', 'Colombia', '@cidenet.com.co');
INSERT INTO public.countries (id, abbreviation, name, domain) VALUES (2, 'USA', 'Estados Unidos', '@cidenet.com.us');
------------------------------------------------------------------------------------------------------------
DROP TYPE IF EXISTS public.status_user cascade ;
create type public.status_user as enum ('enable', 'disable', 'stand-by');
------------------------------------------------------------------------------------------------------------
drop table if exists public.users cascade;
create table if not exists public.users
(
    id                     uuid         not null
        primary key,
    name             varchar(20)  not null,
    others_names           varchar(20),
    last_name        varchar(20)  not null,
    second_last_name       varchar(20)  not null,
    countries_id           integer
        constraint users_countries_fk
            references countries
            on update cascade,
    identification_type_id integer
        constraint users_identification_fk
            references identification_type
            on update cascade,
    identification_number  varchar(20)  not null,
    mail                   varchar(300) not null
        unique,
    admission              date    not null CHECK ( (admission <= current_timestamp)),
    registration           timestamp    not null CHECK (registration >= admission),
    department_id          integer
        constraint users_department_fk
            references department
            on update cascade on delete cascade,
    status                 status_user default 'enable'
);
INSERT INTO public.users(id, name, others_names, last_name, second_last_name, countries_id, identification_type_id, identification_number, mail, admission, registration, department_id, status) VALUES ('2314cdbe-c6c2-4a16-a70b-92c149452eb1', 'Laura', 'Daniela', 'Aguado', 'Rendon', 1, 1, '00010', 'laura.daniela0@correo.com.co', '2022-01-03', '2022-01-12 18:08:05.000000', 4, 'enable');
INSERT INTO public.users (id, name, others_names, last_name, second_last_name, countries_id, identification_type_id, identification_number, mail, admission, registration, department_id, status) VALUES ('a28bb51d-961a-4a09-97b8-90a847627afe', 'Andres', 'Lucca', 'Kennedy', 'Rendon', 1, 1, '1116238356', 'andres.kennedy@correo.com.co', '2022-01-03', '2022-01-12 18:08:05.000000', 4, 'enable');
INSERT INTO public.users (id, name, others_names, last_name, second_last_name, countries_id, identification_type_id, identification_number, mail, admission, registration, department_id, status) VALUES ('0603848a-e623-4e60-b961-383fcb039cb1', 'Aura', '', 'Cadena', 'Rendon', 1, 1, '1116238356', 'aura.cadena@correo.com.co', '2022-01-03', '2022-01-12 18:08:05.000000', 4, 'enable');
INSERT INTO public.users (id, name, others_names, last_name, second_last_name, countries_id, identification_type_id, identification_number, mail, admission, registration, department_id, status) VALUES ('4f8eec07-7bd4-4072-8f19-47b35b9be438', 'Luis', 'Lucca', 'Ocampo', 'Rendon', 2, 1, '1116238356', 'luis.ocampo@correo.com.us', '2022-01-03', '2022-01-12 18:08:05.000000', 4, 'enable');
INSERT INTO public.users (id, name, others_names, last_name, second_last_name, countries_id, identification_type_id, identification_number, mail, admission, registration, department_id, status) VALUES ('68d9e5ac-a147-41e7-a94f-2562cfb951df', 'Carlos', 'Lucca', 'Mendoza', 'Rendon', 1, 1, '1116238356', 'carlos.mendoza@correo.com.co', '2022-01-03', '2022-01-12 18:08:05.000000', 4, 'enable');
INSERT INTO public.users (id, name, others_names, last_name, second_last_name, countries_id, identification_type_id, identification_number, mail, admission, registration, department_id, status) VALUES ('90b37096-2518-4ade-bdc3-f25393e65124', 'Gian', '', 'Lopez', 'Rendon', 2, 1, '1116238356', 'gian.lopez@correo.com.us', '2022-01-03', '2022-01-12 18:08:05.000000', 4, 'enable');
INSERT INTO public.users(id, name, others_names, last_name, second_last_name, countries_id, identification_type_id, identification_number, mail, admission, registration, department_id, status) VALUES ('03fdc57f-e9bf-47d0-a25a-255f4a7f894b', 'Gian', 'Lucca', 'Aguado', 'Rendon', 2, 1, '1116238356', 'gian.aguado@correo.com.us', '2022-01-03', '2022-01-12 18:08:05.000000', 4, 'enable');
INSERT INTO public.users (id, name, others_names, last_name, second_last_name, countries_id, identification_type_id, identification_number, mail, admission, registration, department_id, status) VALUES ('60e1ec05-49c6-4944-807e-afc96d26c4bf', 'Nancy', 'Lucca', 'Parra', 'Ocampo', 1, 1, '1001', 'nancy.parra@correo.com.co', '2022-01-03', '2022-01-12 18:08:05.000000', 4, 'enable');
INSERT INTO public.users (id, name, others_names, last_name, second_last_name, countries_id, identification_type_id, identification_number, mail, admission, registration, department_id, status) VALUES ('c822c1e2-247d-4bb2-83fd-daf82773afd2', 'alberto', 'alfoso', 'de la calle', 'Ruiz', 1, 1, '01', 'alberto.delacalle@correo.com.co', '2022-01-03', '2022-01-12 13:00:00.000000', 4, 'enable');

drop function if exists select_users(varchar, character varying[], character varying[], character varying[], varchar, varchar,
                           integer);
create or replace function select_users(
    search_i varchar, --name , others_names, last_name, second_last_name, identification_number, mail
    countries_i varchar[],
    identifications_types_i varchar[],
    departments_i varchar[],
    status_i varchar,
    cursor_i varchar,
    limit_i int
)
    returns table
            (
                total_rows              int,
                uuid_o                  uuid,
                name_o                  varchar,
                others_names_o          varchar,
                last_name_o             varchar,
                second_last_name_o      varchar,
                country_o               varchar,
                identification_type_o   varchar,
                identification_number_o varchar,
                mail_o                  varchar,
                department_o            varchar,
                status                  status_user
            )
    language plpgsql
as
$$
DECLARE
    i_iterator int = 1;
    i_total    bigint ;
    s_select   varchar;
    s_from     varchar;
    s_where    varchar;
    s_sql      varchar;

BEGIN


    /*
     select * from select_users(
        '',
        array [''],
        array [''],
        array [''],
        'enable',
        '(''Gian'',''03fdc57f-e9bf-47d0-a25a-255f4a7f894b'')',
        2
    );
     */


    s_from := ' FROM users e, countries c, identification_type i, department d ';
    s_where := ' WHERE  e.countries_id = c.id AND e.identification_type_id = i.id AND e.department_id = d.id ';


    -- FILTRO DE BUSQUEDA -------------- BY --> name , others_names, last_name, second_last_name, identification_number, mail
    if (search_i != '') then
        s_where := s_where
                       || ' AND ( '
                       || '  lower(e.name) LIKE ' || chr(39) || '%' || lower(search_i) || '%' || chr(39)
                       || ' OR lower(e.others_names) LIKE ' || chr(39) || '%' || lower(search_i) || '%' || chr(39)
                       || ' OR  lower(e.last_name) LIKE ' || chr(39) || '%' || lower(search_i) || '%' || chr(39)
                       || ' OR  lower(e.second_last_name) LIKE ' || chr(39) || '%' || lower(search_i) || '%' || chr(39)
                       || ' OR  lower(e.identification_number) LIKE ' || chr(39) || '%' || lower(search_i) || '%' ||
                   chr(39)
                       || ' OR  lower(e.mail) LIKE ' || chr(39) || '%' || lower(search_i) || '%' || chr(39)
            || ' ) ';
    end if;
    -- FIN FILTRO DE BUSQUEDA -------------------------------------


    --  FILTRO DE PAISES -------------------------------------
    -- los arrays empieza su indece en 1
    if array_length(countries_i, 1) > 0 then
        if countries_i[1] != '' then
            s_where := s_where || ' AND (  c.id = ' || chr(39) || countries_i[1] || chr(39) || ' ';
            loop
                i_iterator := i_iterator + 1;
                exit when i_iterator > array_length(countries_i, 1);
                s_where := s_where || ' OR  c.id = ' || chr(39) || countries_i[i_iterator] || chr(39) || ' ';
            end loop;
            s_where := s_where || ' ) ' || ' ';
        end if;
    end if;
    i_iterator := 1;
    --   FIN FILTRO DE PAISES-------------------------------------


    --  FILTRO DE TIPO DE DOCUMENTO -------------------------------------
    if array_length(identifications_types_i, 1) > 0 then
        if identifications_types_i[1] != '' then
            s_where := s_where || ' AND (  i.id = ' || chr(39) || identifications_types_i[1] || chr(39) || ' ';
            loop
                i_iterator := i_iterator + 1;
                exit when i_iterator > array_length(identifications_types_i, 1);
                s_where := s_where || ' OR  i.id = ' || chr(39) || identifications_types_i[i_iterator] || chr(39) ||
                           ' ';
            end loop;
            s_where := s_where || ' ) ' || ' ';
        end if;
    end if;
    i_iterator := 1;
    --   FIN FILTRO DE DOCUMENTO-------------------------------------


    --  FILTRO DE DEPARTAMENTOS -------------------------------------
    if array_length(departments_i, 1) > 0 then
        if departments_i[1] != '' then
            s_where := s_where || ' AND (  d.id = ' || chr(39) || departments_i[1] || chr(39) || ' ';
            loop
                i_iterator := i_iterator + 1;
                exit when i_iterator > array_length(departments_i, 1);
                s_where := s_where || ' OR  d.id = ' || chr(39) || departments_i[i_iterator] || chr(39) || ' ';
            end loop;
            s_where := s_where || ' ) ' || ' ';
        end if;
    end if;
    i_iterator := 1;
    --   FIN FILTRO DE DEPARTAMENTOS -------------------------------------


    --   FIN FILTRO ESTADO
    if status_i != '' then
        s_where := s_where || ' AND  e.status  = ' || chr(39) || status_i || chr(39);
    end if;


    -- COUNT TOTAL
    execute ' SELECT count(*) ' || s_from || s_where into i_total;

    s_select := ' SELECT ' || i_total::varchar || '::int as total , e.id, e.name,e.others_names,e.last_name,e.second_last_name,c.name,i.name,' ||
                ' e.identification_number,e.mail,d.name,e.status';

    -- PAGINACION   -----------------------------
    if cursor_i != '' then
        s_where := s_where || ' AND  (e.name, e.id)  >  ' || cursor_i;
    end if;
    -- FIN PAGINACION   -----------------------------


    if limit_i > 0 then
        s_sql := s_select || s_from || s_where || ' ORDER BY e.name, e.id LIMIT ' || limit_i::varchar;
    else
        s_sql := s_select || s_from || s_where || ' ORDER BY e.name, e.id ' ;
    end if;

    -- raise notice 'QUERY -> %', s_sql;
    RETURN QUERY EXECUTE s_sql;

EXCEPTION
    WHEN others THEN
        ROLLBACK;
        RAISE EXCEPTION
            USING ERRCODE = sqlstate
                ,MESSAGE = 'select_users() [' || sqlstate || '] : ' || sqlerrm;
END
$$;
