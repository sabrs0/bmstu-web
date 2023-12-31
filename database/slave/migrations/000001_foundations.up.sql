--
-- PostgreSQL database dump
--

-- Dumped from database version 14.9 (Ubuntu 14.9-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.9 (Ubuntu 14.9-0ubuntu0.22.04.1)

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

--
-- Name: delete_all_foundrisings_of_found(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.delete_all_foundrisings_of_found() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
begin
		DELETE FROM foundrising_tab WHERE found_id = OLD.id;
		IF NOT FOUND THEN 
			RETURN NULL; 
		END IF;
		DELETE FROM transaction_tab 
		WHERE from_id = OLD.id AND from_essence_type = true;
		IF NOT FOUND THEN 
			RETURN NULL; 
		END IF;
		DELETE FROM transaction_tab 
		WHERE to_id = OLD.id AND to_essence_type = false;
		IF NOT FOUND THEN 
			RETURN NULL; 
		END IF;
	RETURN OLD;
end;
$$;


ALTER FUNCTION public.delete_all_foundrisings_of_found() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: foundation_tab; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.foundation_tab (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    password character varying(100) NOT NULL,
    cur_foudrising_amount integer,
    closed_foundrising_amount integer,
    fund_balance real,
    income_history real,
    outcome_history real,
    volunteer_amount integer,
    country character varying(100),
    login character varying(100) NOT NULL
);


ALTER TABLE public.foundation_tab OWNER TO postgres;

--
-- Name: foundrising_tab; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.foundrising_tab (
    id integer NOT NULL,
    found_id integer NOT NULL,
    description text,
    required_sum real,
    current_sum real,
    philantrops_amount integer,
    creation_date date,
    closing_date date
);


ALTER TABLE public.foundrising_tab OWNER TO postgres;

--
-- Name: transaction_tab; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transaction_tab (
    id integer NOT NULL,
    from_essence_type boolean,
    from_id integer NOT NULL,
    to_essence_type boolean,
    sum_of_money real,
    comment text,
    to_id integer NOT NULL
);


ALTER TABLE public.transaction_tab OWNER TO postgres;

--
-- Name: user_tab; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_tab (
    id integer NOT NULL,
    login character varying(100) NOT NULL,
    password character varying(100) NOT NULL,
    balance real,
    charity_sum real,
    closed_fing_amount integer
);


ALTER TABLE public.user_tab OWNER TO postgres;

--
-- Data for Name: foundation_tab; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.foundation_tab (id, name, password, cur_foudrising_amount, closed_foundrising_amount, fund_balance, income_history, outcome_history, volunteer_amount, country, login) FROM stdin;
1	F1	123	0	0	0	0	0	0	Россия	F1
2			0	0	0	0	0	0		
3			0	0	0	0	0	0		
4			0	0	0	0	0	0		
5			0	0	0	0	0	0		
6	F2	12345	0	0	0	0	0	0	США	f1234
7	F3	12345	0	0	0	0	0	0	США	f1235
8	F22	12345	0	0	0	0	0	0	США	f2234
\.


--
-- Data for Name: foundrising_tab; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.foundrising_tab (id, found_id, description, required_sum, current_sum, philantrops_amount, creation_date, closing_date) FROM stdin;
\.


--
-- Data for Name: transaction_tab; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.transaction_tab (id, from_essence_type, from_id, to_essence_type, sum_of_money, comment, to_id) FROM stdin;
1	f	0	f	0		0
2	f	0	f	0		0
3	f	0	f	0		0
4	f	0	f	0		0
\.


--
-- Data for Name: user_tab; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.user_tab (id, login, password, balance, charity_sum, closed_fing_amount) FROM stdin;
1	gduggen0	hbsvo9r	50000	0	0
2	ecane1	JEqZ7pO8xTb	50000	0	0
3	kszymanowicz2	YemtNZNWx	50000	0	0
4	lwaggatt3	uWgosGrjb	50000	0	0
5	clacer4	OcU99nVZb	50000	0	0
6	ihemphrey5	iGdbzOmqO	50000	0	0
7	ffairpo6	INhxdH4WB	50000	0	0
8	cbrevitt7	kCOF36ww	50000	0	0
9	fmcgrory8	zzQKMNY8	50000	0	0
10	hclinkard9	t3rYPAPj5X	50000	0	0
11	cmarryatta	ekrz2LXYlw18	50000	0	0
12	cgeorgievb	wqFuiE0	50000	0	0
13	tjarrattc	5Lkp6Nf	50000	0	0
14	bsapsforded	7DVV6cHq	50000	0	0
15	rfitzgeorgee	jEBJz3zyz	50000	0	0
16	lringsf	vseWtYCuq	50000	0	0
17	lmosbyg	NaIay6jGA	50000	0	0
18	dpomfreyh	cJApLSbp	50000	0	0
19	fbracknalli	ZpsXqqmC	50000	0	0
20	mvanoortj	U0yKBf	50000	0	0
21	pfindenk	LC03I2i	50000	0	0
22	rkobschl	FWi2bIG	50000	0	0
23	mkubelkam	zXTey891He	50000	0	0
24	mboasn	5TP4eJSrG	50000	0	0
25	npillingero	WuG2y4u	50000	0	0
26	tedlynp	48Nv5VZsG	50000	0	0
27	nyurygynq	Mb8jcV	50000	0	0
28	fdelortr	0BI61t4XV	50000	0	0
29	dmatiaseks	8kyRR7BMr	50000	0	0
30	rashbeet	u0Hag8twA	50000	0	0
31	aducarneu	z2URVMWiv	50000	0	0
32	gburnhillv	PNrqQhVWldIT	50000	0	0
33	berdesw	3I0PRYaCq7	50000	0	0
34	llorenx	J7oV0wbZ2xu	50000	0	0
35	hdebretty	ditgQw1	50000	0	0
36	taronowiczz	tEGXYZLLK	50000	0	0
37	gmennithorp10	fEQT8oN	50000	0	0
38	cripping11	NIw3KzCnrE85	50000	0	0
39	ofreyne12	K08ZlVJr71	50000	0	0
40	eklimpke13	nJFPIU1	50000	0	0
41	rfranzini14	2a8bNjISVfrn	50000	0	0
42	cledstone15	BLsG6FLz94Yb	50000	0	0
43	pdewsbury16	gGizZq0S	50000	0	0
44	shatt17	MV4nYg	50000	0	0
45	gmudge18	O7yI0l	50000	0	0
46	lshegog19	qXs81Iq	50000	0	0
47	cserjent1a	sKRhydFj6GvJ	50000	0	0
48	fgarnar1b	9VSlrX	50000	0	0
49	nshelley1c	Psj6eH4oD	50000	0	0
50	meastmond1d	RECu13sQ	50000	0	0
51	edrissell1e	ukHlElxg	50000	0	0
52	mflello1f	EX4uppf4Z	50000	0	0
53	nshingles1g	kh3SqT2zmx	50000	0	0
54	mcollins1h	beu2mdjXQNCu	50000	0	0
55	elithcow1i	L4vOxXVxcvgx	50000	0	0
56	acarrivick1j	yGeNu0psc	50000	0	0
57	ndouch1k	MtkJ1n	50000	0	0
58	ccondict1l	JKWLoRqiUEhx	50000	0	0
59	adiamant1m	KR6wDNMfP	50000	0	0
60	vtommeo1n	g7vH9lY	50000	0	0
61	bpetow1o	6wXEpjO69Ogc	50000	0	0
62	arowlands1p	CUqCiji1wb	50000	0	0
63	gfrenchum1q	R5O3ioFbjM4o	50000	0	0
64	jhopfner1r	gkBmXkvc6HW	50000	0	0
65	ebellsham1s	HO7ysazJ	50000	0	0
66	elanegran1t	c5BacC1eY	50000	0	0
67	aboyford1u	xHaRFGzyzZx	50000	0	0
68	etomaszynski1v	CEzSxcY	50000	0	0
69	groocroft1w	jOlV3ZAna5	50000	0	0
70	kpolkinhorn1x	hhkEWDBp7p	50000	0	0
71	apayne1y	mtV6oWCGE	50000	0	0
72	akilminster1z	zZMJDzSQ87lA	50000	0	0
73	afurlong20	GneupHL	50000	0	0
74	sverdon21	ygkoazRk	50000	0	0
75	igodden22	A1hDIVk	50000	0	0
76	adalliwatr23	L29vv4cXr3Q5	50000	0	0
77	cbeardshaw24	l31aDb	50000	0	0
78	mbrinsford25	LLQBUqqOByX	50000	0	0
79	echipping26	3duSteS	50000	0	0
80	oworshall27	SO1NZJ	50000	0	0
81	jrobson28	4ahzXDKUxyT	50000	0	0
82	dshropshire29	2IkXJ969Ax	50000	0	0
83	tcaldera2a	dR4QTI2	50000	0	0
84	mjados2b	9j1CYVSEMQWp	50000	0	0
85	tcaldicot2c	zuU1fs4ekA	50000	0	0
86	lnoller2d	5WZ6ll	50000	0	0
87	tdymidowski2e	cTmBOrO	50000	0	0
88	nlissimore2f	nSEDMe	50000	0	0
89	ptourville2g	KXdkWBS	50000	0	0
90	kmcmanamon2h	RECyt88Xo	50000	0	0
91	rmarcussen2i	4tOyFGkoVLGl	50000	0	0
92	vbeves2j	3vgoWTvT	50000	0	0
93	ojakolevitch2k	MuCIZqLRYBTe	50000	0	0
94	edrance2l	5RVPf4Xw	50000	0	0
95	acorkill2m	x86TvuFkVg	50000	0	0
96	caddington2n	TjktgdS	50000	0	0
97	dhidderley2o	dRC4w3Zp	50000	0	0
98	tmulvany2p	WbaHKPbVtm5	50000	0	0
99	shampton2q	a50pSqXkKqjq	50000	0	0
100	bmeriott2r	N1GMe1Iuzc6	50000	0	0
101	rwoodburne2s	GX2F2j0	50000	0	0
102	mmarson2t	amHUD03yNY	50000	0	0
103	vcaps2u	akjc7pdNOK	50000	0	0
104	badriani2v	73UxPM6	50000	0	0
105	prau2w	9p2xM3bdf7b	50000	0	0
106	drawle2x	q3sHcK4Hi	50000	0	0
107	elecornu2y	6qCrm6ATP	50000	0	0
108	yjossel2z	FG9ACc7Ct	50000	0	0
109	ialman30	kWQ69LNArami	50000	0	0
110	ftillman31	AOhwq9woP9ut	50000	0	0
111	pmacsween32	BonWq1	50000	0	0
112	ecrosse33	khH95BuS	50000	0	0
113	chalhead34	JmUMSJ	50000	0	0
114	adoppler35	iRVfVt	50000	0	0
115	ablainey36	K9YYMn	50000	0	0
116	ehug37	FNld4uZy	50000	0	0
117	asayer38	YB7sr86ad	50000	0	0
118	pgonnard39	03KrKoBtw4K	50000	0	0
119	jnolan3a	ONEwTj	50000	0	0
120	sysson3b	G9yl1guhnZ	50000	0	0
121	rmorcom3c	F92T1GOa	50000	0	0
122	chilbourne3d	2SmNCNDwRZZ4	50000	0	0
123	gbrazener3e	a0fJubBJ	50000	0	0
124	lmaestrini3f	tO9boOPYU5	50000	0	0
125	rdulanty3g	pgi5ID4J87pQ	50000	0	0
126	sbaison3h	lA2ceeN5G1g	50000	0	0
127	gwrotchford3i	0UTC8grPlP	50000	0	0
128	mcayet3j	tqwytk	50000	0	0
129	mnewey3k	5KzRekD	50000	0	0
130	rboskell3l	vjA1D2AqOAtV	50000	0	0
131	jogborne3m	X10fMV	50000	0	0
132	jsulland3n	kMO4WMs	50000	0	0
133	cnisco3o	PiZRry07X	50000	0	0
134	hblyden3p	UrcYZ2B	50000	0	0
135	celgie3q	dZ6FBv7W	50000	0	0
136	gthorsby3r	zpnc8Xp1J8y	50000	0	0
137	hemerson3s	lE3jGm	50000	0	0
138	ggaudon3t	m4WaXsjQxtB	50000	0	0
139	ikeegan3u	JOXcb7UM1	50000	0	0
140	fgentsch3v	5VTOtWSyP	50000	0	0
141	jkilcullen3w	KqoZ0ZTK87bV	50000	0	0
142	jdare3x	5ucGoN6N	50000	0	0
143	hmacsporran3y	hvyWgHBIi	50000	0	0
144	bquickfall3z	3Ijoa8NzY7n	50000	0	0
145	jborrill40	HY8rUFXhJq	50000	0	0
146	acaro41	iPqkJLy	50000	0	0
147	fmilvarnie42	fr3vmI	50000	0	0
148	mdarker43	uyRrZqlu	50000	0	0
149	afahy44	NEUwpwcTNi	50000	0	0
150	eleband45	ZDr9XoRDxqvx	50000	0	0
151	glehrmann46	ZKenAEO42n	50000	0	0
152	ncrockatt47	mQHKaxdCjzk	50000	0	0
153	rgroarty48	JbR9j50	50000	0	0
154	jdilley49	O2GBepc0	50000	0	0
155	tswindall4a	AItY7oR	50000	0	0
156	jraecroft4b	eb6BQuB7ybJR	50000	0	0
157	chaddacks4c	Ybxi7S	50000	0	0
158	pevered4d	WDUrcPyTHSKo	50000	0	0
159	aferney4e	7kcMlQa1zY	50000	0	0
160	emussen4f	5hULOkAW	50000	0	0
161	tsyvret4g	mj03Kei	50000	0	0
162	blovekin4h	Rc6NIaVkj	50000	0	0
163	rkulas4i	c0jNCmQ	50000	0	0
164	ibayle4j	5AigciSRk	50000	0	0
165	passiratti4k	OaT22r	50000	0	0
166	mmaciaszczyk4l	eMuFv9RXpD11	50000	0	0
167	thandscomb4m	D0mDQTN	50000	0	0
168	clawley4n	r8TLcwTIG	50000	0	0
169	idowden4o	euCBekUlyoG8	50000	0	0
170	skerrane4p	cr7Y31H	50000	0	0
171	mdottrell4q	iQxC9g84Ap	50000	0	0
172	gjentges4r	TEhSb9Z	50000	0	0
173	lcaller4s	VvdU353Kx	50000	0	0
174	ppickle4t	qsB1b3z6	50000	0	0
175	jleadbetter4u	y6kim7hlTC	50000	0	0
176	pchamley4v	TRvpYN5NXboq	50000	0	0
177	jfellgatt4w	uQTr7Bl5	50000	0	0
178	rcarrett4x	SCXQA5g13F	50000	0	0
179	sstreat4y	HAP3TE1	50000	0	0
180	tteffrey4z	RWvHrx0J	50000	0	0
181	spalethorpe50	G8lzOb	50000	0	0
182	murling51	BHy9cfob0	50000	0	0
183	cduffield52	8iAcIoZ	50000	0	0
184	nclipson53	HdMAwq9Np4	50000	0	0
185	bdurkin54	EHMg0VT	50000	0	0
186	gmarkel55	rMxQJjH	50000	0	0
187	rpirelli56	lgGkIZ	50000	0	0
188	smuddicliffe57	vLe4Ah	50000	0	0
189	csolland58	nRZsfMOEAX	50000	0	0
190	cchristou59	CF1YLaN	50000	0	0
191	sellett5a	9Kle973FCr	50000	0	0
192	ydedomenico5b	xgpSHd	50000	0	0
193	dmacneely5c	bXJQaF	50000	0	0
194	ahaskur5d	rhlDwm3M	50000	0	0
195	cchattelaine5e	6G83eW	50000	0	0
196	bbeavan5f	Iit675f1t	50000	0	0
197	dsaile5g	yweQpxZe	50000	0	0
198	sdarco5h	Ap5shB6UL	50000	0	0
199	epuddefoot5i	Bzk9KG	50000	0	0
200	gchesterman5j	g4ribK5GN	50000	0	0
201	rboatman5k	ldmu3vQi3IP	50000	0	0
202	gbunston5l	HpvaIhtlp4XY	50000	0	0
203	vforst5m	tbBEB0N2	50000	0	0
204	lkesper5n	RSHo7DT	50000	0	0
205	lvaneev5o	vB9lNDcnm	50000	0	0
206	lsynan5p	uEcjWhy3nNbx	50000	0	0
207	etyhurst5q	vp5U0huKxiXT	50000	0	0
208	dstidworthy5r	025er4Ns	50000	0	0
209	nsoffe5s	TP1wR5rZY	50000	0	0
210	btoseland5t	LQXH2ErVn2IB	50000	0	0
211	jlindback5u	aqU1IGa	50000	0	0
212	ntedorenko5v	3SivGc	50000	0	0
213	iglavias5w	Hlevv4qrv	50000	0	0
214	scamlin5x	zd1XLcl9CV9	50000	0	0
215	meddowis5y	Cxf2eVWSHScL	50000	0	0
216	jbaglan5z	2FOaFY	50000	0	0
217	ccordaroy60	741ScAeYi	50000	0	0
218	bwildman61	K07pJ9LC5A	50000	0	0
219	gsharpe62	TMbc4E4AE	50000	0	0
220	cdilon63	LOc6SeAxBpS	50000	0	0
221	latwool64	ZffJgL6	50000	0	0
222	eabbots65	q5OysWq	50000	0	0
223	dtuckie66	NK5uRyAfAHaK	50000	0	0
224	psiviter67	7De7kBDmETu7	50000	0	0
225	wbiddell68	k5wY9NMT9q	50000	0	0
226	dsaul69	IPTTESdWSs	50000	0	0
227	mdalziell6a	C9O7Tzt	50000	0	0
228	dlamberton6b	v5PadHpfz	50000	0	0
229	etrigwell6c	Yt2GjN5Xhz	50000	0	0
230	abrumpton6d	6u6qwo	50000	0	0
231	lfarrimond6e	SgTpxRFbu4Ym	50000	0	0
232	jworg6f	zHhM1m	50000	0	0
233	tthomann6g	1LPkMKR	50000	0	0
234	zmoehle6h	J2Tw5QN	50000	0	0
235	hswindells6i	qeF1qia5s	50000	0	0
236	cizachik6j	POKb2y	50000	0	0
237	bbailes6k	Dk7yCaNgw	50000	0	0
238	jtirone6l	tiQMGy7B	50000	0	0
239	glapley6m	hrWGIMMP	50000	0	0
240	heastway6n	QmC9zsI2zpqf	50000	0	0
241	rsidebotham6o	e9wt20m0	50000	0	0
242	mtomasz6p	fn6dU7x0	50000	0	0
243	rstannus6q	RLn8a6YDMgc1	50000	0	0
244	rbaldazzi6r	rb7acnizsw	50000	0	0
245	bberni6s	3SrvI66jS16V	50000	0	0
246	mballchin6t	Tco4kYPw	50000	0	0
247	ghurren6u	34Zhhiqp8WNI	50000	0	0
248	cfardon6v	vdVPiXvYw4K	50000	0	0
249	cpoulter6w	EwhugGfG44	50000	0	0
250	mhaddrill6x	LoZ19QpR5J	50000	0	0
251	ebengle6y	SNzOv2	50000	0	0
252	sleupold6z	vDl3PvLMbW	50000	0	0
253	speppett70	yoMc8O	50000	0	0
254	dbotwright71	PWwj7TkV	50000	0	0
255	prebanks72	jMUOjS2Im	50000	0	0
256	cgladdish73	alcZpHZm	50000	0	0
257	rfitzgeorge74	Agap9mmD	50000	0	0
258	jbrisard75	CbhZFjKbLz	50000	0	0
259	gdezamora76	dKhNfji7EMm	50000	0	0
260	cmessam77	E1EKwuMz4u	50000	0	0
261	csibbson78	szJAPhqROsH	50000	0	0
262	nchristofor79	MzHBVFCJwy	50000	0	0
263	mburrows7a	0GvXN5g93	50000	0	0
264	dcummine7b	4oE8Se	50000	0	0
265	gtunaclift7c	lYjD6GRq	50000	0	0
266	tperrin7d	PCwnGwnMhYfB	50000	0	0
267	bpapez7e	leRG47Sk	50000	0	0
268	gderill7f	tDDlWoF5U	50000	0	0
269	raviss7g	pmeUArvOroH	50000	0	0
270	yparkins7h	85gBtH	50000	0	0
271	gwandrach7i	MzC1Rcb	50000	0	0
272	mquinnette7j	u0Z3TlN9	50000	0	0
273	hhurndall7k	w7YD4kTPUiMY	50000	0	0
274	vweavers7l	qQcnLGBESG9l	50000	0	0
275	mhoulworth7m	aPfJvYbyj	50000	0	0
276	cbevens7n	jWOe4cOxJ	50000	0	0
277	afloyde7o	oMW1ebeUxY	50000	0	0
278	iboycott7p	01HDfeubkZ	50000	0	0
279	lwestphalen7q	iJCoUTI	50000	0	0
280	lstansby7r	kQCDmT41Kdcv	50000	0	0
281	wphilippet7s	kndzcY0yH5K8	50000	0	0
282	bandrosik7t	8oj0gkl	50000	0	0
283	ahempshall7u	SSUFKAju81	50000	0	0
284	dbosnell7v	iH1bivI1uQDc	50000	0	0
285	aswanbourne7w	2GfZze	50000	0	0
286	iiskowicz7x	5sfJwYIdN	50000	0	0
287	rkimbell7y	iptb1lUK	50000	0	0
288	bturner7z	rKUlh3TxCr	50000	0	0
289	adealmeida80	m0mnWbxg3g	50000	0	0
290	dwhatley81	4AXibbLum	50000	0	0
291	jingliss82	bQgbPhN3K	50000	0	0
292	agroven83	gSlwkNeR7PIF	50000	0	0
293	kmcdougall84	tZB0hDQ	50000	0	0
294	jpryke85	mOrCvLEBiPT	50000	0	0
295	lmorphey86	llGNywIQDN	50000	0	0
296	tludlamme87	lxWyRb1e9	50000	0	0
297	lpetrowsky88	BSkLlQfVxy5	50000	0	0
298	ctarborn89	ClBAPI	50000	0	0
299	pchazette8a	H3OjYXNcd9	50000	0	0
300	dbew8b	ll86E3KKDQ	50000	0	0
301	lkiendl8c	0NRNEDvP	50000	0	0
302	gpodbury8d	dRhtsLK8XL	50000	0	0
303	bcorradetti8e	YGqyWFY26ITa	50000	0	0
304	hgiacomozzo8f	ulxmqLQmYDrN	50000	0	0
305	smulchrone8g	7pcizGQVAE	50000	0	0
306	lhoff8h	XrZXNLZFUQ	50000	0	0
307	bospellissey8i	KXRa153bENm	50000	0	0
308	tbroxap8j	GgLcbk0HlBS	50000	0	0
309	cwaterhous8k	W1IhVX	50000	0	0
310	bsimonich8l	dN5uv9yU1k	50000	0	0
311	mprinett8m	IbvmUiJsbOwb	50000	0	0
312	mburchill8n	ytygWxFkz7F	50000	0	0
313	egodehardsf8o	JPCUnNi6qmZ	50000	0	0
314	narlow8p	XWxQ1xiC	50000	0	0
315	fmcgeffen8q	nDgxFvuOp2	50000	0	0
316	mruggiero8r	9JDbfd1h	50000	0	0
317	abaudet8s	YoLyRP9J0t	50000	0	0
318	ygermain8t	WxibUNQ4PxBL	50000	0	0
319	smapletoft8u	86bk78b	50000	0	0
320	tmckeney8v	32d6uQQM	50000	0	0
321	nmitskevich8w	SYGWvm	50000	0	0
322	kbignal8x	DPrEAfI1TYK	50000	0	0
323	lannis8y	wVVfRePx	50000	0	0
324	glindsell8z	UdsWPK	50000	0	0
325	elemmanbie90	u5fjCP1yTUN	50000	0	0
326	mlafontaine91	5YAb9aROZ	50000	0	0
327	mpluthero92	f9J5C1rULdcn	50000	0	0
328	ewigg93	SrEcLz2Ghq	50000	0	0
329	mrimour94	dn2O323	50000	0	0
330	lwillshear95	KyZZRKoZ9	50000	0	0
331	cmatias96	kok0p3LhyYaR	50000	0	0
332	nfeeny97	ZZSj32PrOo1	50000	0	0
333	mstenner98	ZRDiaC3ZUMLH	50000	0	0
334	floudwell99	aDIGqL	50000	0	0
335	dcaw9a	rF6cqJm	50000	0	0
336	kleverich9b	2mGYDe	50000	0	0
337	gweerdenburg9c	x9p0paJ0gEYA	50000	0	0
338	cmiddell9d	LsOMxf2Q	50000	0	0
339	hcaller9e	0zbanQ0f	50000	0	0
340	jbrudenell9f	3SDFR2yr4zoU	50000	0	0
341	kreiglar9g	J54P2x	50000	0	0
342	zsay9h	WE4VwNAFfebT	50000	0	0
343	barchambault9i	vWdBjeR	50000	0	0
344	jhaggith9j	nabaMU6j8	50000	0	0
345	sskirving9k	6JFF5HsgG3e	50000	0	0
346	rlayton9l	nGhpokuSoE9	50000	0	0
347	detheredge9m	ZKNY5zIHu	50000	0	0
348	dnorthin9n	r5pwbZR0Yi	50000	0	0
349	cwellstood9o	gffNd0sd	50000	0	0
350	scarson9p	prMpirfQEImL	50000	0	0
351	grobertz9q	1gfN9XDoWd	50000	0	0
352	nwitz9r	qaAUbWk	50000	0	0
353	istonham9s	t34nxH	50000	0	0
354	krappaport9t	WJmyA4	50000	0	0
355	vianitti9u	4cXkXH	50000	0	0
356	mfarady9v	AV5CHw	50000	0	0
357	dfidian9w	Y0vO5yFpJK3Z	50000	0	0
358	sconduit9x	HB4r7uZ	50000	0	0
359	slockitt9y	O12VKIWp	50000	0	0
360	tlarkings9z	zM5nWU2JF9	50000	0	0
361	kcurnowa0	MClY0pJf3n	50000	0	0
362	rfrankea1	xTgUcU7t5Da	50000	0	0
363	hramarda2	zmha8K8Bl	50000	0	0
364	bfancea3	LlE3BvV2lM	50000	0	0
365	mcaunta4	pvEY9B6CE2T	50000	0	0
366	hnorgatea5	xWE1t897	50000	0	0
367	mbreckona6	QMuzRekzZFc	50000	0	0
368	jplowmana7	TUbCLL9	50000	0	0
369	estearnsa8	1KozhojH2	50000	0	0
370	hbittena9	LbbLeRlm	50000	0	0
371	rchoppingaa	0HCg0bVx4	50000	0	0
372	cbickellab	GqxeVRj	50000	0	0
373	hstroobantac	U13zZzBn9o5	50000	0	0
374	alindnerad	DsBZ3c	50000	0	0
375	htyttertonae	N1M5GgX3G	50000	0	0
376	peudallaf	ea7hgH7	50000	0	0
377	lstoteag	fOYqEqlIdP	50000	0	0
378	lhumphrysah	eg85Uc2oq7	50000	0	0
379	hpiffordai	PledrRqH1njI	50000	0	0
380	gverdeyaj	2a6sdc4mxGVQ	50000	0	0
381	lrosellak	E1rTcB	50000	0	0
382	rwyseal	YehapFjkHWc	50000	0	0
383	lbrabanam	USGGhnz	50000	0	0
384	acowlingan	VfzyM6wGnX	50000	0	0
385	drichardetao	gaZOEL8QIJ	50000	0	0
386	pmoonap	I7R03hD	50000	0	0
387	nsetteraq	WQBvdGQsaJS	50000	0	0
388	cmacphailar	t9hFhYfXVR	50000	0	0
389	ademullettas	J9QNkL	50000	0	0
390	imcilriachat	Vgt0g2	50000	0	0
391	djostanau	Cw15vXwM	50000	0	0
392	jsousaav	zIR7wemI	50000	0	0
393	cgiriardelliaw	cxT4MsfO6N	50000	0	0
394	edearnlyax	TC2DRY30XJR5	50000	0	0
395	jhebbleay	rfCAtf	50000	0	0
396	dhuskaz	J0OA00aUkY	50000	0	0
397	nchewb0	PcT0v5	50000	0	0
398	echalfainb1	MTFiCRl	50000	0	0
399	adreweb2	U10XUoFwm5x	50000	0	0
400	btattershawb3	n8kVNnl2	50000	0	0
401	rforrestillb4	GXvNo8	50000	0	0
402	cbellhouseb5	PMm9OMaMd0	50000	0	0
403	bfenwickb6	ln31kCaEDk	50000	0	0
404	hworswickb7	6SydHu8CIO	50000	0	0
405	chadrillb8	RDyjBa0	50000	0	0
406	lboeckeb9	L9RIGTAt	50000	0	0
407	kkasparba	4suC7IUog4uj	50000	0	0
408	malbertsbb	AaDlc5Qh	50000	0	0
409	gslymbc	rbgaxjT	50000	0	0
410	aburrellsbd	VFKMAVWza	50000	0	0
411	ppeploebe	Lb2vMd	50000	0	0
412	aescalerabf	pgxblFtojQUw	50000	0	0
413	rknottonbg	lpsTqgKpLo0	50000	0	0
414	jkidsleybh	UKynva7yaUJ8	50000	0	0
415	tbrestonbi	vuxCJObN	50000	0	0
416	hcoveneybj	OTHVjPfE	50000	0	0
417	avaheybk	h69lNdH7gPM	50000	0	0
418	jkiosselbl	oSCNjVQaL6uv	50000	0	0
419	jleneybm	CA6lUsC3p5	50000	0	0
420	asauratbn	hA7Qakd	50000	0	0
421	hnisiusbo	AcDK2D	50000	0	0
422	jromandbp	ipZ2tAgw	50000	0	0
423	wlittlemorebq	fDBxJChg3F	50000	0	0
424	cdigginsbr	Ds5F3Oq9v	50000	0	0
425	aeimbs	Jctbxt2Ude	50000	0	0
426	cbeauvaisbt	PsJJXHqCHqW9	50000	0	0
427	twadlybu	7S9AwKrkfWc	50000	0	0
428	pnaultybv	92KZvxceWa	50000	0	0
429	ekeywoodbw	cFNzYV	50000	0	0
430	lbertomeubx	KieUBVz	50000	0	0
431	vnormanvilleby	wegFcwg	50000	0	0
432	dtzarbz	HAIvyI2	50000	0	0
433	memmsc0	dHrmLaD	50000	0	0
434	rpayleyc1	EZA5w0RD	50000	0	0
435	dbaggallyc2	N5VWOU	50000	0	0
436	tselbiec3	rzlPhAtw7c	50000	0	0
437	hstannusc4	hts1RGf7	50000	0	0
438	egaitonec5	kgIvqoOODT9e	50000	0	0
439	rtobiasc6	lqgsivhaz	50000	0	0
440	jllorensc7	OJfeDd	50000	0	0
441	nagarc8	tWlWvVIxe78P	50000	0	0
442	ctollitc9	asUpMN8J	50000	0	0
443	asharplingca	p9P7uH3	50000	0	0
444	nyaldencb	uzOhwzl2j56k	50000	0	0
445	shackfordcc	MaYLqKZ	50000	0	0
446	lredmondcd	JddtxpA	50000	0	0
447	cpehrssonce	2zcsb6Qwf1ER	50000	0	0
448	abiagionicf	yN1Ji0mKb5	50000	0	0
449	dskoylescg	9f9cdVxY0bM	50000	0	0
450	dcampesch	Aqsdn3m2tTn	50000	0	0
451	jorgillci	DEz4uv	50000	0	0
452	lcronkshawcj	35LCH2HaJX	50000	0	0
453	senstoneck	TsVf7Wc	50000	0	0
454	plongeacl	Sb4RO64C	50000	0	0
455	apauluzzicm	SPUErIEHBTcB	50000	0	0
456	jbourleycn	sogODmp3Xw	50000	0	0
457	bleederco	sykXYu	50000	0	0
458	cbattincp	drqHBw	50000	0	0
459	egubbincq	8rTsw9YWMxxm	50000	0	0
460	gabbiscr	zngt06lZk	50000	0	0
461	fwindleycs	rgPEpL	50000	0	0
462	cfatscherct	Gs9t822rkr	50000	0	0
463	egeckscu	f61umzb7oBxs	50000	0	0
464	kcanetcv	68crTevqqV	50000	0	0
465	jseatoncw	wmlZbMikECqY	50000	0	0
466	edeeprosecx	HqTZc0cYnm9W	50000	0	0
467	hcoveycy	JAICe6Jq4c1n	50000	0	0
468	ckepecz	JP6LYc1wrd	50000	0	0
469	awilksd0	Qes7HlT	50000	0	0
470	nchristofold1	tQ1egaH	50000	0	0
471	dfishlockd2	4oPZZ1	50000	0	0
472	kdowd3	UEois7	50000	0	0
473	tpaddeleyd4	tXoLBXx	50000	0	0
474	ddowsingd5	4wM5WsSG	50000	0	0
475	dslimmingd6	MXeJt4Kdrm2	50000	0	0
476	akalinked7	GRGDeAA	50000	0	0
477	cchampkinsd8	AwIcrmxPSQt	50000	0	0
478	mmacnaughtond9	WbcuAMnR	50000	0	0
479	bgoodreadda	p1ByCBWgHS	50000	0	0
480	mgoldbourndb	1eL8Q30LhgEe	50000	0	0
481	fcurleydc	XwIzkRHl	50000	0	0
482	gdampierdd	0IJKyQve	50000	0	0
483	saxstonde	bVDNtiOg3B	50000	0	0
484	dscurreydf	HQw1PNzSlBFH	50000	0	0
485	shaughandg	XfpqFYNp	50000	0	0
486	psotworthdh	nC2qZmD0zZ	50000	0	0
487	itackledi	uwMMpVFQP	50000	0	0
488	pborlanddj	uaZTgi3TQz	50000	0	0
489	tjacobsohndk	k7fFOl1	50000	0	0
490	adowtrydl	EeMCrdUt3jZ	50000	0	0
491	mspiveydm	gBeIk8H	50000	0	0
492	esunnersdn	H3g5VT4tqkm	50000	0	0
493	scristofolodo	dEomDN5GF	50000	0	0
494	dellerdp	ikoq3E	50000	0	0
495	lservantdq	nWLwVTCkG4MT	50000	0	0
496	kferrydr	lk8zlXToZ	50000	0	0
497	bwissbyds	vlDp0I	50000	0	0
498	dgozneydt	7aAYpKyJqn	50000	0	0
499	bsheltondu	7dFJssx6L	50000	0	0
500	rharberdv	b2iZ3ds0Uh37	50000	0	0
501	meastcottdw	fSLxfaegwJi	50000	0	0
502	nwebburndx	TwVDceZre	50000	0	0
503	clindldy	pZBxuIDp	50000	0	0
504	bhailstondz	ELQe3RCcRa	50000	0	0
505	aelwarde0	2vJRXgX	50000	0	0
506	nedele1	zz6XVX	50000	0	0
507	mmorricee2	F1M0d8	50000	0	0
508	derwine3	3JRIp2Z	50000	0	0
509	tsimsone4	BxdX1DIFCT	50000	0	0
510	bworshalle5	5cZzZaMQLtP4	50000	0	0
511	rjollimane6	O633Nhyyd	50000	0	0
512	mclemette7	xMISkN5	50000	0	0
513	lkwieteke8	Lo7cl6vZ7E	50000	0	0
514	sdantone9	rB2BDd52Oq	50000	0	0
515	bevinsea	mTb73EiD	50000	0	0
516	ebartolomeazzieb	h6tMUZ	50000	0	0
517	fdadgeec	KTpYuph	50000	0	0
518	calthroped	BUPqHs	50000	0	0
519	coakmanee	BBRCHEV	50000	0	0
520	jpreddleef	rZhLNsbBU	50000	0	0
521	beagleneg	IDGD9g	50000	0	0
522	wyakovliveh	7H5r8Kk8	50000	0	0
523	jsikorskiei	JPUb6ZVH	50000	0	0
524	fgutej	FSx7V2hCowS	50000	0	0
525	haleavyek	J6QMjnarPAy	50000	0	0
526	lkingcoteel	Ta3veRF9	50000	0	0
527	cvasserem	DNvmCNLpivan	50000	0	0
528	sbrookeren	WW6vYTE	50000	0	0
529	gwornumeo	nJvyUnVVk	50000	0	0
530	ipummellep	tOpXGmx	50000	0	0
531	kaldricheq	84DcMn4	50000	0	0
532	gmccooker	7Ax1uv	50000	0	0
533	vdixceyes	wtyPe9Or	50000	0	0
534	hlarteret	PvnrzuZ58Gg	50000	0	0
535	lbroxholmeeu	vILuOS7	50000	0	0
536	abroschkeev	e58ilF2pP	50000	0	0
537	fjoderliew	WQDX6YHLrXp	50000	0	0
538	jthonex	YMMmoUW1L3kF	50000	0	0
539	etroppmanney	Hu32rimwwB	50000	0	0
540	ssimonuttiez	Xh2P9bZHk	50000	0	0
541	dhaggasf0	wDvSJomPI03	50000	0	0
542	tkellowayf1	KTgUVa	50000	0	0
543	dveillardf2	wTYizoPzmR3	50000	0	0
544	bjambrozekf3	mMNxec7ao	50000	0	0
545	plipprosef4	KhK2WmcpK	50000	0	0
546	vtighef5	Aszgh686c1pF	50000	0	0
547	odunsterf6	Hro5NVwYolR	50000	0	0
548	bcahnf7	3afxLl	50000	0	0
549	kziemkef8	hVSxySC	50000	0	0
550	fbehninckf9	8vTVe18tRLf	50000	0	0
551	shedanfa	Jq0mFU	50000	0	0
552	tmaddocksfb	THAXP6nFO	50000	0	0
553	kbruneaufc	vpm7Ii	50000	0	0
554	zclifffd	XjJrQuospQp	50000	0	0
555	llillegardfe	xfTAXFbB	50000	0	0
556	tselfff	fksFJrX	50000	0	0
557	dpickinfg	PPUfdajlnwAs	50000	0	0
558	wvaldesfh	avvZYzVq9	50000	0	0
559	mcahalinfi	drSTjCN0Y	50000	0	0
560	aronischfj	Ys90RBJl	50000	0	0
561	ciacovinofk	jp5ZcAVsK	50000	0	0
562	coliphantfl	t5RPwfzA2Q	50000	0	0
563	ahavefm	7YgwpL9	50000	0	0
564	mfordycefn	5dp4CGUKw	50000	0	0
565	gyeardsleyfo	U3tTuMIuOu6	50000	0	0
566	lwilksfp	Vp0f1JPVsA	50000	0	0
567	bpickoverfq	4vxl62Y0ZkFh	50000	0	0
568	obraderfr	4NwLPMj	50000	0	0
569	nmelonbyfs	EpdcTXmXEB	50000	0	0
570	cradboneft	94blpIpLCOR	50000	0	0
571	bgoldbergfu	GaIpBDl66o	50000	0	0
572	lleakfv	i8MpGN	50000	0	0
573	mlegricefw	VutPjiWS5p	50000	0	0
574	wvossefx	bPnNE4cTuQK	50000	0	0
575	ctorresify	cMudNBQWm6kn	50000	0	0
576	eseysfz	v53HPvcyQUDf	50000	0	0
577	tspringthorpeg0	PSa6N91T	50000	0	0
578	cjoscelyng1	ljK0Fsv5L	50000	0	0
579	kmcgonnellg2	VriDf1IR	50000	0	0
580	mdietscheg3	2YXKPhzmy4	50000	0	0
581	uhalligang4	jQ9n9b0Nk5	50000	0	0
582	rureng5	AlXnfsG04Qb	50000	0	0
583	cvarndellg6	hXXbRgfuf	50000	0	0
584	sjanuszg7	HjUxXJver	50000	0	0
585	sboulterg8	4OcZaGlDPN	50000	0	0
586	tlomaxg9	uDCo1AlgwYE5	50000	0	0
587	wfulkga	YKBD2VzH4Il	50000	0	0
588	aruprechtgb	klDMOEcaQ	50000	0	0
589	hizhakgc	eh5IkF1OB	50000	0	0
590	hmeconigd	Tqag3iFKww	50000	0	0
591	wbrownbillge	IGXV7TB5V	50000	0	0
592	mbeushawgf	eggVcz153la	50000	0	0
593	asandellsgg	1pvzwq9oQHN	50000	0	0
594	syitshakgh	J0lGda	50000	0	0
595	cloveridgegi	C6jyQ6ogHlzE	50000	0	0
596	ndabnotgj	CVknNM	50000	0	0
597	hdisleygk	BmjVMNXZOi	50000	0	0
598	jfrankisgl	2knd5F	50000	0	0
599	tdinzeygm	iwQ01kQ7h	50000	0	0
600	esalatinogn	wZHfr0yFFFbl	50000	0	0
601	heggersgo	70kB4lFiBsY	50000	0	0
602	soliveiragp	8AxhIw	50000	0	0
603	abartomieugq	x9BFSM	50000	0	0
604	fsignorelligr	PVqd35Dki5J	50000	0	0
605	hprettjohngs	eNoU7P0	50000	0	0
606	nfeehangt	uygfhjVYj	50000	0	0
607	gbroadnickegu	4orp4Esau15	50000	0	0
608	mwindressgv	QaxGeAQYKa2	50000	0	0
609	dhatchergw	Jfjl3KXF	50000	0	0
610	kduminigx	LI5cajW0q0	50000	0	0
611	ulocklessgy	2rx4FGiInrw3	50000	0	0
612	lmcturleygz	8QXAlNHaei	50000	0	0
613	adelleschih0	6w24r08	50000	0	0
614	rbyrkmyrh1	mMJAGQf2	50000	0	0
615	pmulhillh2	RNO54FKNm	50000	0	0
616	bomaileyh3	78G72JFqT	50000	0	0
617	jcallowh4	jocBDtX	50000	0	0
618	ejeckellh5	kToWqPhbwMLG	50000	0	0
619	sissacovh6	5bFpdTiz4t	50000	0	0
620	tturfush7	hJL7xHdLTj	50000	0	0
621	kabrianih8	DZbgDuXn0	50000	0	0
622	vhaglintonh9	PUwZcoLIt867	50000	0	0
623	aallmondha	baEml53M	50000	0	0
624	cbenoisthb	T28hCngQ	50000	0	0
625	mcasinahc	YUMSZc6Ru5	50000	0	0
626	kbouslerhd	j7zvWq6M1l	50000	0	0
627	sarnaudihe	QgczjJw	50000	0	0
628	zavenhf	RguTY5	50000	0	0
629	msopphg	8G3U6XgO4cG	50000	0	0
630	mmaidstonehh	QxdXl2izHDYs	50000	0	0
631	cbillhamhi	Ndbcww	50000	0	0
632	cpettingallhj	vf0ONQQ4CBt	50000	0	0
633	glongleyhk	5wPfR9b	50000	0	0
634	pvanderweedenburghl	GdsgkwST	50000	0	0
635	dmullalyhm	vRc2pD44T8rv	50000	0	0
636	oclifforthhn	TkYBdhkxMxH	50000	0	0
637	wgreenoho	lkX6XN4qJU4	50000	0	0
638	ehowsanhp	NXkAQTX	50000	0	0
639	gculleyhq	kwiJUwrl	50000	0	0
640	acleaveshr	tBt8FanbjYu	50000	0	0
641	kholbornhs	LwHowjdZ9B	50000	0	0
642	cpirelliht	dkERsvms	50000	0	0
643	cflynnhu	2cDlnTZLwe	50000	0	0
644	aharraginhv	jSU40tm0H	50000	0	0
645	mcollumhw	wHFzkTdl1FA	50000	0	0
646	msherwynhx	DISQ6LEII0oW	50000	0	0
647	rmilanhy	X0Y0VDNTPsW	50000	0	0
648	ckennifeckhz	fK2In7nuqivm	50000	0	0
649	stembletti0	TwMknoO	50000	0	0
650	vpinchbecki1	834ovvwHlfx	50000	0	0
651	lhopferi2	LPVPNV	50000	0	0
652	tstefii3	65TwXvi	50000	0	0
653	kriddingtoni4	U4AqLiXH8	50000	0	0
654	bkindoni5	gQ0gnnh	50000	0	0
655	bgaskarthi6	YdH6OK	50000	0	0
656	rcaghani7	X7CxdEFqDfn	50000	0	0
657	cstannahi8	6MlgIiXaELJ	50000	0	0
658	agreastyi9	WlkMg5	50000	0	0
659	ksimoesia	zGsd4uFJz	50000	0	0
660	jissacoffib	agF4YMPllN	50000	0	0
661	tgallawayic	D7Tu6pzRZz	50000	0	0
662	smccarlieid	dr9G2pOm343	50000	0	0
663	jhuikerbyie	XW3FKXQFvW	50000	0	0
664	llonghurstif	gixh24	50000	0	0
665	awithringtonig	UlUOsf	50000	0	0
666	jwarnesih	z0h5zQ	50000	0	0
667	nnapoleonii	CM5ioHnHR4DC	50000	0	0
668	csizeij	kt5Hfc5Vw4	50000	0	0
669	khebnerik	AAJJ5pxJ3D	50000	0	0
670	rjillingsil	pppGuHacPH1	50000	0	0
671	dwoodingtonim	69zCcLNa	50000	0	0
672	imegaineyin	wiAWtpCq	50000	0	0
673	gscaddonio	5Jnuznz	50000	0	0
674	astrakerip	PY2lOoMa	50000	0	0
675	fceschiiq	0zYyJyxKbzN	50000	0	0
676	rbrabenir	zKjkT59Q	50000	0	0
677	cmccomiskieis	UMXYxTpZC	50000	0	0
678	cstoppit	opve7A	50000	0	0
679	ademaineiu	lKSkuo	50000	0	0
680	rlasselleiv	mApry0Ju	50000	0	0
681	mwimpeneyiw	wmaLvnxm7Z	50000	0	0
682	dnortonix	2YHjGTg41GA	50000	0	0
683	lfaltiniy	OlpgZis6UWN	50000	0	0
684	gingreyiz	7KfFimY	50000	0	0
685	mormesherj0	gMELyvINM	50000	0	0
686	eluisj1	qSMbH2grU3gD	50000	0	0
687	rdecourtj2	KkIADK	50000	0	0
688	bdrancej3	JRw3yaDzW	50000	0	0
689	dgaudinj4	Sftv1r7cciX	50000	0	0
690	pgordonj5	FEq5JNMcr	50000	0	0
691	mbertolinj6	b0SKNk	50000	0	0
692	dmcentagartj7	XmjEI4ofMP7	50000	0	0
693	rdurdlej8	Too0QYJgTcd	50000	0	0
694	dcampbelldunlopj9	tQgzKyK	50000	0	0
695	jvasquezja	6OsddRJI	50000	0	0
696	jbackhurstjb	xmRNjx	50000	0	0
697	cbatesjc	D1hAa3bq	50000	0	0
698	msievejd	ZwaoA7zXL	50000	0	0
699	daudenisje	JSGo45aQT	50000	0	0
700	allewellenjf	99ImGWDdcrYB	50000	0	0
701	gvanrossjg	H6obN17GCU	50000	0	0
702	wpearmanjh	QeHPqknkDj	50000	0	0
703	lmacgibbonji	KSh7agClF	50000	0	0
704	spardonjj	M80DPx	50000	0	0
705	fmajorjk	1TCe1sUWnA	50000	0	0
706	tslixbyjl	3ER6H8a5uev	50000	0	0
707	eeddowisjm	b1TsiQe	50000	0	0
708	athurlbournejn	VuczuhjAv	50000	0	0
709	agierthjo	hdsFdtx	50000	0	0
710	tnuttingjp	OuQ2VRLKpi	50000	0	0
711	mwoakesjq	7uEZ74p	50000	0	0
712	twinsiowieckijr	Ind20K3zdU	50000	0	0
713	eolyonovjs	Rb1H88i	50000	0	0
714	sbessonjt	bYDTF3umGqz7	50000	0	0
715	ispacieju	Eg0mo2bk1	50000	0	0
716	lburchmorejv	FPfJPFQB0G	50000	0	0
717	spizziejw	np0Zr0NqQ5V	50000	0	0
718	dlangtryjx	0QQY5eKN	50000	0	0
719	mespinheirajy	zpPxzoi	50000	0	0
720	dfaughnanjz	LnTNdimnycOW	50000	0	0
721	dpettik0	xKh1Zyj	50000	0	0
722	ainklek1	FfuIXg2xg43R	50000	0	0
723	wsipsonk2	HM0NglR6ql	50000	0	0
724	rludeek3	GJQMuOL7	50000	0	0
725	kbernakiewiczk4	xvw1M941x	50000	0	0
726	wheballk5	VnAYeohhIHJh	50000	0	0
727	csmeethk6	igoCXgren	50000	0	0
728	ahoutbyk7	x4JhkTZ	50000	0	0
729	dputleyk8	PKtXoMT9	50000	0	0
730	jroblouk9	85ZZsDh2	50000	0	0
731	svittetka	A1EOTP	50000	0	0
732	fbottinikb	KoIMklQpZTPV	50000	0	0
733	slawrensonkc	ZkhWUQuth6WE	50000	0	0
734	nmcduallkd	7slP3qp5k	50000	0	0
735	lmactimpanyke	1LmwGZFjOQnH	50000	0	0
736	cgoodbarkf	Cquv9x80H7Q	50000	0	0
737	kspratlingkg	Qy2GRMaZh8	50000	0	0
738	nleckenbykh	VH0HXg	50000	0	0
739	cslefordki	EdOqgJijEFw	50000	0	0
740	afernelykj	EgJsy8Bt	50000	0	0
741	dlowensohnkk	mD6zyA5euCXX	50000	0	0
742	gbeckekl	DeGPdK	50000	0	0
743	dmandreskm	JFzwEB0	50000	0	0
744	csalvadorikn	XmAjiKf	50000	0	0
745	mloudwellko	fai2DK8fHgVj	50000	0	0
746	jhebblewhitekp	ZSWCI1pQ3v	50000	0	0
747	fscoynekq	ul8D3MCYk	50000	0	0
748	mmillomkr	qQsTp7	50000	0	0
749	mrushmerks	uXXwAzNQ	50000	0	0
750	lblackburnkt	iSfGZt	50000	0	0
751	dnenciku	GhzaXoafs	50000	0	0
752	mdungeykv	KjKBM3Ro	50000	0	0
753	estubbskw	IHAUddyBx	50000	0	0
754	jburgherkx	YwgT1Rq4bG	50000	0	0
755	pdillingstonky	uOXy2M999Q3	50000	0	0
756	cshakesbykz	X2EPPpfc4h	50000	0	0
757	astannersl0	y75aArHmBnd	50000	0	0
758	blambdinl1	rPh4rNO	50000	0	0
759	rdefrainel2	bXzalHjBRz	50000	0	0
760	djellisl3	Kux73MTw9NZ7	50000	0	0
761	cbonnerl4	8jVM27	50000	0	0
762	bbootherl5	6s6Wfd	50000	0	0
763	abutterwickl6	bnwTfzVyJn	50000	0	0
764	ssalll7	Oueha6	50000	0	0
765	jsutherbyl8	FvfTWgXBfBy	50000	0	0
766	sbastonl9	vPIXRQfX	50000	0	0
767	gkestertonla	Na6ery0P	50000	0	0
768	mellereylb	q2qdMy	50000	0	0
769	lcorainilc	ztd1RP	50000	0	0
770	pmackowleld	XhkpzuG8	50000	0	0
771	kcarnegyle	bO5xFlEjA	50000	0	0
772	dpaxfordelf	bqPBOzW1ig	50000	0	0
773	sparkhouselg	wHlQacJn8	50000	0	0
774	dpeschkelh	OSYbrSXpYy0	50000	0	0
775	strevillionli	fRpb3zmAY5pu	50000	0	0
776	tsheppardlj	A851ujlE3	50000	0	0
777	athreadgilllk	Ro3vXu1GXOV	50000	0	0
778	lworsamll	LKIkbvYLeb2	50000	0	0
779	esabinlm	3vLTTSiG	50000	0	0
780	mbonnesenln	WTsChgZn6Tcb	50000	0	0
781	fbetjeslo	6DmMlY5qL50	50000	0	0
782	dlongstreethlp	cP1NV6TW9l	50000	0	0
783	uflockhartlq	tgDrKo35	50000	0	0
784	jbodycombelr	KlmgCwp1	50000	0	0
785	rceeleyls	coyYgoPKP6WB	50000	0	0
786	tnicklinlt	ZHgt7P	50000	0	0
787	dselewaylu	szh7hcQ	50000	0	0
788	adorbinlv	heSwnx6	50000	0	0
789	csmizlw	EBRPQkIY	50000	0	0
790	ahollowslx	fHLX6yFx7KTJ	50000	0	0
791	bpoxtonly	FHAuQS	50000	0	0
792	wcrampsylz	TUsZM1ML	50000	0	0
793	rgerholzm0	gYeISMGi	50000	0	0
794	platheym1	lbLm9C158	50000	0	0
795	dbeckenhamm2	hhujd2PEW	50000	0	0
796	tlysaghtm3	8EwHw5	50000	0	0
797	ctonnesenm4	LDuz0ebP	50000	0	0
798	ashimonim5	EagXUG	50000	0	0
799	vgeraschm6	RZrj7vDqG	50000	0	0
800	ccasseym7	G7F4Te6sKSb	50000	0	0
801	atomensonm8	WXT0bhX	50000	0	0
802	hshurmerm9	7uFFX6uWUNl	50000	0	0
803	olovejoyma	4CdYkVPXYn3	50000	0	0
804	bhandscombmb	evTIth	50000	0	0
805	psneakermc	1mOPWTO	50000	0	0
806	dmatelaitismd	teyY35Jm	50000	0	0
807	spaffotme	n2IK44H	50000	0	0
808	mfaulksmf	jjL3yxaMmLH	50000	0	0
809	mteligamg	yaUQuDqvtNC0	50000	0	0
810	cgristwoodmh	XMKmzrdLLB	50000	0	0
811	brobkenmi	x0HxVBbt	50000	0	0
812	lstrafenmj	AjIV9gLAF	50000	0	0
813	cfrancklynmk	nRKN3rsyo	50000	0	0
814	mwhetnellml	i0jjaCJvHQ7j	50000	0	0
815	fshevelsmm	1quA7sbyzKkt	50000	0	0
816	mpeertmn	ADiQ5P	50000	0	0
817	jgitthousemo	ffukyVxSG	50000	0	0
818	mcollissonmp	FwTYUGCQQOU	50000	0	0
819	cstebbingsmq	EtAjMybss3E5	50000	0	0
820	eeldenmr	kty9n9ZG5mp	50000	0	0
821	dscorahms	EDdiHNh	50000	0	0
822	ppaulleymt	tK8pnRyNN	50000	0	0
823	ncreevymu	2W9ZhHpZUFST	50000	0	0
824	jbrennenmv	wmXsA4TvD	50000	0	0
825	sjouanotmw	6dcd80HLkiKB	50000	0	0
826	bpfaffemx	0eZ9JJRnKvxG	50000	0	0
827	ksnellmanmy	REiSSW46	50000	0	0
828	gnorridgemz	DlpYfPGsG	50000	0	0
829	mgoakesn0	OYFEP0	50000	0	0
830	mgerrann1	0FZfO39ViV	50000	0	0
831	dminillon2	JsuFFXV	50000	0	0
832	twrenn3	9VbQ2lgK	50000	0	0
833	weisenbergn4	pxkNl1AAw5	50000	0	0
834	avreibergn5	fC63im0	50000	0	0
835	dhehnken6	ynmXj1N	50000	0	0
836	ekinneallyn7	l15elpUZeC4N	50000	0	0
837	dcowhign8	Vk9VOa	50000	0	0
838	dolenechann9	mWFC14Ncrg8L	50000	0	0
839	cwarrenerna	gPXvamMV	50000	0	0
840	jhinrichnb	TYvTU4e	50000	0	0
841	rforstallnc	znOysOY9ipNY	50000	0	0
842	jharrildnd	WlGYni	50000	0	0
843	agaiterne	awalGDylx	50000	0	0
844	koliveynf	xbJl3MyJs1Yz	50000	0	0
845	ashirdng	gdxJr2	50000	0	0
846	jchartresnh	Z1QUJ11	50000	0	0
847	mdehooghni	JMhabFMbL4S	50000	0	0
848	abinleynj	yP0od6	50000	0	0
849	hmcawnk	02HoF8i4FoQ	50000	0	0
850	mcrennannl	oUCLIkwl1LSp	50000	0	0
851	jstandbridgenm	JDtivtNdNhC	50000	0	0
852	jgeorgelnn	9HybJ6p8i44A	50000	0	0
853	mbrambleyno	OnTTbbIdw	50000	0	0
854	bhuncotenp	ZqNdkcsK	50000	0	0
855	rcuberleynq	VkKxiaI2hUZ	50000	0	0
856	droycroftnr	3ryFWxCuql	50000	0	0
857	boldalens	5rm50NPcbh	50000	0	0
858	hharriskinent	ZglqFbNroc	50000	0	0
859	npatershallnu	gbRDTdhiZLDM	50000	0	0
860	emcilraithnv	HRyXuabXszwu	50000	0	0
861	rgoodaynw	Uy4Esh	50000	0	0
862	azappelnx	mETKt4q5AzA	50000	0	0
863	rgrevilleny	5kNqHJ8E	50000	0	0
864	wscarlonnz	WDIiQb	50000	0	0
865	fsousao0	LK287B	50000	0	0
866	hbrownseyo1	0gLBrnaIY	50000	0	0
867	kfibbingso2	60FS78vA	50000	0	0
868	bbliveno3	X0sNQh3r7Bc	50000	0	0
869	krillstoneo4	zWy7Mh	50000	0	0
870	moveyo5	sGodbI	50000	0	0
871	gdoftyo6	l7g2tLDooHG	50000	0	0
872	amcpeeterso7	cw2IAHBmLyft	50000	0	0
873	driveo8	OgfYp7u3O8	50000	0	0
874	bbowskillo9	2UQh8gy3E	50000	0	0
875	amalinoa	g8zx27ns	50000	0	0
876	nfriedlosob	avsrl2j1F	50000	0	0
877	mdutnelloc	Se4cdON5HeS	50000	0	0
878	mmclugishod	RdIXfY9	50000	0	0
879	mbouldonoe	NsCcZHRXw	50000	0	0
880	malleryof	GbDiXbs	50000	0	0
881	wczajkaog	uiRxVT	50000	0	0
882	agaishoh	bbdRNwDvq2j	50000	0	0
883	mosoriooi	vwCQuzzXsIk	50000	0	0
884	hcopinoj	MqWGZBKt	50000	0	0
885	otamlettok	BG1wMO	50000	0	0
886	pduhamelol	V3kCvuA	50000	0	0
887	wwordsworthom	mVISG5Sr	50000	0	0
888	kmaccheyneon	Pcnj2k	50000	0	0
889	ljakubovitsoo	2zx2Psj	50000	0	0
890	mgerattop	T4lQbdkXoGzJ	50000	0	0
891	kbelfrageoq	lDh9CXY9l0vy	50000	0	0
892	gavrasinor	zHnw2Ox	50000	0	0
893	jdicksonos	LTaco0	50000	0	0
894	jgourleyot	srirnWi	50000	0	0
895	bmishou	plfmT2MEE	50000	0	0
896	aespadaterov	qi5CJP	50000	0	0
897	gmargettow	DdsRuUt	50000	0	0
898	vaugustineox	FyPjyLmJs	50000	0	0
899	rwymanoy	FTNSCqmH73	50000	0	0
900	omaccreeoz	TI4Zl5LJNQQ	50000	0	0
901	ahebbesp0	wFaQYvn	50000	0	0
902	amussettip1	yxbMvah8gdgt	50000	0	0
903	mspourep2	qedy08eh	50000	0	0
904	qgleavep3	AwDQoDIXLYWC	50000	0	0
905	hmustchinp4	chh8wxIzy	50000	0	0
906	ccailep5	LnM4oDQGg	50000	0	0
907	gmaccaigp6	rDRMkxDIcLH	50000	0	0
908	swakelinp7	bpAKYX	50000	0	0
909	mmoffattp8	9dUw3ugbh	50000	0	0
910	gcancottp9	oNjChnVJuKt	50000	0	0
911	umenduspa	eTBvhb2	50000	0	0
912	eleavoldpb	JiMHeOECOoj	50000	0	0
913	qpremblepc	iV4MPJ	50000	0	0
914	maccombepd	sziBLjyDM	50000	0	0
915	cspilstedpe	ft9CkW	50000	0	0
916	asilkstonepf	nKNrH6xNZuG	50000	0	0
917	hgarnsworthpg	yGP1xyxklhNu	50000	0	0
918	mchippph	0GOyfB	50000	0	0
919	gimpettpi	TnFQ7vsSj	50000	0	0
920	ttrevorpj	dQjRWWsQ	50000	0	0
921	mwildborepk	15Krw7YHD	50000	0	0
922	hhammertonpl	RJiOasUer	50000	0	0
923	rminettpm	4hkyxke77aSp	50000	0	0
924	mtorrespn	RrNvR00p	50000	0	0
925	tmckimmpo	Vo9mOO	50000	0	0
926	obattertonpp	QBhiYfDd	50000	0	0
927	ljimenezpq	xfiKV1QbLbQr	50000	0	0
928	zblasiakpr	bsYej1ae	50000	0	0
929	tcromwellps	6mwzPxU	50000	0	0
930	lcoullpt	JTabheUHhN	50000	0	0
931	mkleinholzpu	E84xQZxaR6P	50000	0	0
932	tspearespv	lsEIgZ	50000	0	0
933	aullettpw	1A4P4ps3oi	50000	0	0
934	hsivernspx	lydzKw	50000	0	0
935	bsavatierpy	s6wzYwz94zPh	50000	0	0
936	byaldrenpz	lrm3tZg5y8CE	50000	0	0
937	troglieroq0	6YDLMJecFzE	50000	0	0
938	apresnellq1	zRqKUH	50000	0	0
939	fpeplerq2	MIcoLDM	50000	0	0
940	fdormanq3	u9mqUL7	50000	0	0
941	acantonq4	fAhFJo	50000	0	0
942	bfoulgerq5	5Sjdw7gchSX	50000	0	0
943	ymertgenq6	xDcDpLA91dn	50000	0	0
944	bfilleryq7	mmN4BMfQyIW	50000	0	0
945	tclorleyq8	xVMywSUwHmB	50000	0	0
946	agratrexq9	aNKB2V	50000	0	0
947	kbeabyqa	muCbfr5RwKT	50000	0	0
948	llorentzenqb	Og7yYa8y	50000	0	0
949	fdowthwaiteqc	uAGtjWB5TKD	50000	0	0
950	ahardageqd	EYwIvB	50000	0	0
951	mhackingeqe	OUXydCaLq	50000	0	0
952	mwiggqf	wcvwhB89DqPc	50000	0	0
953	tsibilleqg	s1viD9nM	50000	0	0
954	dmulhallqh	SzFJt7McW2Q	50000	0	0
955	wfeatherqi	ByiH0iKp	50000	0	0
956	rpetrovqj	MprGoKynn	50000	0	0
957	kinggallqk	ajSRWfXVGpg	50000	0	0
958	bfentonql	pEXsV98Nvn8r	50000	0	0
959	nkaesmakersqm	jdkGxR4HQ	50000	0	0
960	ksydenhamqn	AwTZkX7NrZC	50000	0	0
961	dgrinstedqo	O30Xyk	50000	0	0
962	gandreiniqp	HEwfWx	50000	0	0
963	msommertonqq	TPugJz7gmDK	50000	0	0
964	rleaheyqr	wgKtDTCJuyS	50000	0	0
965	agerkensqs	nv0QyV5wNZ	50000	0	0
966	swilhelmyqt	u2Axdl1DXCP	50000	0	0
967	sstandingfordqu	e0kuELk0	50000	0	0
968	egoldbergqv	HD9yE1M8R	50000	0	0
969	pbarhamsqw	PsaobC206R	50000	0	0
970	gdyeqx	9kmrxK3j2M	50000	0	0
971	bgreenhillqy	FSZK0xARk5nR	50000	0	0
972	byellowleyqz	I7W4oGdJ	50000	0	0
973	bloddenr0	kZn78Xs	50000	0	0
974	ncochr1	rfdNuUMGO	50000	0	0
975	gdisleyr2	KX3MZLLtwH	50000	0	0
976	dmeierr3	7vbXPN5JI	50000	0	0
977	pfishlockr4	5rPWEmlSgMN5	50000	0	0
978	zlewinsr5	DYo2aRmUs3f	50000	0	0
979	ekohrsr6	RdHaF7vzW	50000	0	0
980	vgrcicr7	pY8PhzpGHa	50000	0	0
981	dmathieur8	x4D6FRH3qYq	50000	0	0
982	mbleslir9	3WySo0o	50000	0	0
983	dmclarnonra	ceDXqd	50000	0	0
984	barneckerb	jCIO68kmaIG	50000	0	0
985	lcuxsonrc	YpH0LkDqKllY	50000	0	0
986	kbattrickrd	17AGMAUAeIIz	50000	0	0
987	rfavellre	Uy9xeM	50000	0	0
988	awemmrf	gsInWDVjvxMZ	50000	0	0
989	jgunthorperg	25dabXvaSk	50000	0	0
990	asouthonrh	BKfxz0Q0JpY	50000	0	0
991	mderobertori	sXCE5FIuiiO	50000	0	0
992	mwareingrj	EgeBKJ2UCvX9	50000	0	0
993	amatticcirk	ey30gfGw2V	50000	0	0
994	prolfsrl	JLysXc	50000	0	0
995	bosbaldestonrm	5BCAVvqmEPJ	50000	0	0
996	cgoodacrern	5TjT2F2	50000	0	0
997	mconnerryro	YqTnbcBRMhxi	50000	0	0
998	lbuckamrp	JgN8UVX	50000	0	0
999	caxworthyrq	KURQihkg	50000	0	0
1001			0	0	0
1002			0	0	0
1003			0	0	0
1004			0	0	0
1005	string	string	0	0	0
\.


--
-- Name: foundation_tab foundation_tab_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.foundation_tab
    ADD CONSTRAINT foundation_tab_pkey PRIMARY KEY (id);


--
-- Name: foundrising_tab foundrising_tab_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.foundrising_tab
    ADD CONSTRAINT foundrising_tab_pkey PRIMARY KEY (id);


--
-- Name: transaction_tab transaction_tab_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction_tab
    ADD CONSTRAINT transaction_tab_pkey PRIMARY KEY (id);


--
-- Name: user_tab user_tab_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_tab
    ADD CONSTRAINT user_tab_pkey PRIMARY KEY (id);


--
-- Name: foundation_tab foundrisings_deleting; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER foundrisings_deleting AFTER DELETE ON public.foundation_tab FOR EACH ROW EXECUTE FUNCTION public.delete_all_foundrisings_of_found();


--
-- PostgreSQL database dump complete
--

drop user if exists ronly1;
drop user if exists ronly2;
drop role if exists foundations_readonly;
create role foundations_readonly  WITH LOGIN PASSWORD 'ronly'; --LOGIN;
grant connect on database foundations to foundations_readonly;
grant usage on schema public to foundations_readonly;

grant select on all tables in schema public to foundations_readonly;
