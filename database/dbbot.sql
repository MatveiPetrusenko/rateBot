--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5 (Ubuntu 14.5-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.5 (Ubuntu 14.5-0ubuntu0.22.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: easy_economic_exercises; Type: TABLE; Schema: public; Owner: bot_admin
--

CREATE TABLE public.easy_economic_exercises (
    exercise_id integer NOT NULL,
    data_exercise character varying(1024) NOT NULL,
    answer character varying(64) NOT NULL,
    score_point integer NOT NULL
);


ALTER TABLE public.easy_economic_exercises OWNER TO bot_admin;

--
-- Name: easy_economic_exercises_exercise_id_seq; Type: SEQUENCE; Schema: public; Owner: bot_admin
--

CREATE SEQUENCE public.easy_economic_exercises_exercise_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.easy_economic_exercises_exercise_id_seq OWNER TO bot_admin;

--
-- Name: easy_economic_exercises_exercise_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: bot_admin
--

ALTER SEQUENCE public.easy_economic_exercises_exercise_id_seq OWNED BY public.easy_economic_exercises.exercise_id;


--
-- Name: medium_economic_exercises; Type: TABLE; Schema: public; Owner: bot_admin
--

CREATE TABLE public.medium_economic_exercises (
    exercise_id integer NOT NULL,
    data_exercise character varying(1024) NOT NULL,
    answer character varying(64) NOT NULL,
    score_point integer NOT NULL
);


ALTER TABLE public.medium_economic_exercises OWNER TO bot_admin;

--
-- Name: medium_economic_exercises_exercise_id_seq; Type: SEQUENCE; Schema: public; Owner: bot_admin
--

CREATE SEQUENCE public.medium_economic_exercises_exercise_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.medium_economic_exercises_exercise_id_seq OWNER TO bot_admin;

--
-- Name: medium_economic_exercises_exercise_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: bot_admin
--

ALTER SEQUENCE public.medium_economic_exercises_exercise_id_seq OWNED BY public.medium_economic_exercises.exercise_id;


--
-- Name: user_results; Type: TABLE; Schema: public; Owner: bot_admin
--

CREATE TABLE public.user_results (
    user_id integer NOT NULL,
    user_name character varying(512) NOT NULL,
    score_point integer NOT NULL,
    progress integer
);


ALTER TABLE public.user_results OWNER TO bot_admin;

--
-- Name: easy_economic_exercises exercise_id; Type: DEFAULT; Schema: public; Owner: bot_admin
--

ALTER TABLE ONLY public.easy_economic_exercises ALTER COLUMN exercise_id SET DEFAULT nextval('public.easy_economic_exercises_exercise_id_seq'::regclass);


--
-- Name: medium_economic_exercises exercise_id; Type: DEFAULT; Schema: public; Owner: bot_admin
--

ALTER TABLE ONLY public.medium_economic_exercises ALTER COLUMN exercise_id SET DEFAULT nextval('public.medium_economic_exercises_exercise_id_seq'::regclass);


--
-- Data for Name: easy_economic_exercises; Type: TABLE DATA; Schema: public; Owner: bot_admin
--

COPY public.easy_economic_exercises (exercise_id, data_exercise, answer, score_point) FROM stdin;
1	Хватит ли величины вклада, равной 1000 д.е.,\nположенной сегодня в банк под 10%, для\nтого чтобы через 10 лет внести плату за\nобучение, равную 2500 д.е.	Да	5
2	Банк выдает кредит 300 000 руб. на 3 года под\n10% годовых.\nОпределите сумму, которую придется вернуть\nзаемщику.	399300	5
3	Четыре года назад покупатель приобрел дом.\nОн рассчитал, что его сегодняшняя\nстоимость составляет 207360 долл., зная, что\nежегодно недвижимость дорожала на 20%,\nопределите цену покупки.	100000	5
4	От эксплуатации автомобиля предприятие может\nполучать в течение 4 лет доход в размере 120 у.д.е.\nв год.\nКакую минимальную сумму предприятие должно\nполучить от продажи автомобиля, чтобы в случае\nвложения вырученных денег в банк под 15% на 4\nгода иметь доход не ниже, чем результат от\nэксплуатации автомобиля.\nОстаточная стоимость автомобиля через 4 года будет\nравна 100 у.д.е.	331,6	5
5	Финансовый менеджер предприятия предложил Вам инвестировать 10000 руб. в его предприятие, пообещав возвратить 15000 руб. через 3 года. Какова процентная ставка прибыльности предложенного варианта?	14,47	5
\.


--
-- Data for Name: medium_economic_exercises; Type: TABLE DATA; Schema: public; Owner: bot_admin
--

COPY public.medium_economic_exercises (exercise_id, data_exercise, answer, score_point) FROM stdin;
6	Предположим Вы купили шестилетний 8-ми процентный сберегательный сертификат стоимостью $1,000.  Если проценты начисляются ежегодно, какую сумму Вы получите по окончанию контракта?	1586.87	10
7	Если безрисковая ставка составляет 10%, ожидаемая доходность рынка 20%, «бета»-коэффициент портфеля 0.8, то ожидаемая доходность инвестиционного портфеля составляет?	0.18	10
8	Проведя усовершенствование технологического процесса предприятие в течение пяти последующих лет планирует получение ежегодное увеличение денежного дохода на $10,000. Эти деньги оно собирается немедленно вкладывать под 10 процентов годовых, желая через пять лет накопить сумму для приобретения нового оборудования. Какую сумму денег предприятие получит через пять лет?	61.05	10
9	Пусть инвестиции в проект к концу первого года его реализации составят 20 000 руб. В последующие четыре года ожидаются годовые доходы по проекту: 6 000 руб., 8 200 руб., 12 600 руб., 18 800 руб. Рассчитать чистую текущую стоимость проекта к началу первого года, если процентная ставка составляет 10% годовых.	13216.93	10
10	Инвестор с целью инвестирования рассматривает 2 проекта, рассчитанных на 5 лет. Проекты характеризуются следующими данными:  по 1-му проекту – начальные инвестиции составляют  550 тыс. руб., ожидаемые доходы за 5 лет соответственно 100, 190, 270, 300 и 350 тыс. руб.; по 2-му проекту – начальные инвестиции составляют  650 тыс. руб., ожидаемые доходы за 5 лет соответственно 150, 230, 470, 180 и 320 тыс. руб. Определить, какой проект является наиболее привлекательным для инвестора при ставке банковского процента – 15% годовых и в качестве ответа дать чистую приведенную стоимость	225392.59	15
\.


--
-- Data for Name: user_results; Type: TABLE DATA; Schema: public; Owner: bot_admin
--

COPY public.user_results (user_id, user_name, score_point, progress) FROM stdin;
248410765	@nikitenko_anna	0	0
355364196	@ma1thew	0	0
\.


--
-- Name: easy_economic_exercises_exercise_id_seq; Type: SEQUENCE SET; Schema: public; Owner: bot_admin
--

SELECT pg_catalog.setval('public.easy_economic_exercises_exercise_id_seq', 1, false);


--
-- Name: medium_economic_exercises_exercise_id_seq; Type: SEQUENCE SET; Schema: public; Owner: bot_admin
--

SELECT pg_catalog.setval('public.medium_economic_exercises_exercise_id_seq', 1, false);


--
-- Name: easy_economic_exercises easy_economic_exercises_pkey; Type: CONSTRAINT; Schema: public; Owner: bot_admin
--

ALTER TABLE ONLY public.easy_economic_exercises
    ADD CONSTRAINT easy_economic_exercises_pkey PRIMARY KEY (exercise_id);


--
-- Name: medium_economic_exercises medium_economic_exercises_pkey; Type: CONSTRAINT; Schema: public; Owner: bot_admin
--

ALTER TABLE ONLY public.medium_economic_exercises
    ADD CONSTRAINT medium_economic_exercises_pkey PRIMARY KEY (exercise_id);


--
-- PostgreSQL database dump complete
--

