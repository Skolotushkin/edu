# Глава 2: Язык Go

## Язык Go

### 1. Что вы узнаете из этой главы?
- Происхождение Go: когда и кем оно было создано.
- Мотивы создания Go.
- Каковы основные характеристики языка?

### 2. Рассмотренные технические концепции
- Время сборки
- Статически типизированный язык
- Параллелизм
- Сборщик мусора
- Зависимость от программного обеспечения

### 3. Миф о сотворении мира
Существует миф о создании Go. Язык родился в офисе Google, и это произошло во время очень долгой сборки, которая заняла 45 минут.

Эту историю рассказывает Роб Пайк в [@go-at-google]. Это дает нам ценную информацию о мотивах создания Go. Время сборки было слишком долгим и болезненным... они должны были найти способ избежать их; это было отправной точкой для Go Genesis.

Роберт Гриземер, Кен Томпсон и Роб Пайк — разработчики, которые начали работать над Go еще в 2007 году. Роб Пайк утверждает, что к середине 2008 года язык был «в основном спроектирован, и реализация (компилятор, время выполнения) начала работать». После этого в 2008 году к команде присоединились Иэн Лэнс Тейлор и Расс Кокс [@pike2009go].

Go — это язык программирования с открытым исходным кодом, поддерживаемый сообществом и основной командой разработчиков, работающих в Google. 16 марта 2011 года состоялся первый релиз Go. (Он назывался "r56"). Go версия 1 была выпущена 28 марта 2012 года.

× Бумажное и цифровое издание этой книги доступны [здесь](#).
Я также снял видеокурс по созданию реального проекта с помощью Go.

### 4. Мотивация
Go (или Golang) был создан Google для решения проблем компании. Чтобы лучше понять причины, стоит прочитать основной доклад Роба Пайка [@go-at-google].

С какими проблемами сталкивается программное обеспечение в крупных компаниях по всему миру?

- Кодовая база сервисов Google огромна. У Google миллионы строк кода.
- Эти строки написаны на разных языках: C, C++, Java и других.
- Время сборки этих приложений «растянулось на многие минуты и даже часы».
- Обновление некоторых частей приложения может быть дорогостоящим.

Цель первых сусликов заключалась в том, чтобы облегчить жизнь разработчикам путем:

- Резкое сокращение времени сборки программ.
- Разработка языка, который будет прост в изучении, чтении и отладке для молодых разработчиков, знакомых с C, C++ или Java.
- Разработка эффективной системы управления зависимостями.
- Создание языка, который может создавать программное обеспечение, которое хорошо масштабируется на аппаратном обеспечении.

#### 4.1 Определение некоторых понятий
- **Время сборки**: количество времени, необходимое компилятору для создания машиночитаемого исполняемого файла.
- **Статически типизированный язык**: Давать точное определение этого понятия в настоящее время преждевременно. Мы рассмотрим этот термин в следующих главах.
- **Зависимость**: часть программного обеспечения, которая используется другим программным обеспечением.
- **Масштабируемость**: способность программы обрабатывать растущее количество задач, которые должны быть выполнены. Например, веб-сайт считается масштабируемым, если он может принимать растущее количество запросов без простоев или увеличения задержки загрузки.

### 5. Ключевые особенности Go
Создатели Go сосредоточили свои усилия на нескольких критически важных дизайнерских решениях:

- Компилируемый язык
- С помощью семантики легко понять и освоить
- Статическая типизация
- Встроенный параллелизм, с которым разработчикам легко работать
- Надежное управление зависимостями
- Сборщик мусора

Основная цель, по словам Роба Пайка, заключалась в том, чтобы дать разработчикам простой в освоении язык для «проектирования больших программных проектов».

#### 5.1 Некоторые понятия
- **Параллелизм**: Программа является параллельной, когда задачи могут выполняться не по порядку или в частичном порядке.
- **Сборщик мусора (часто называемый GC)**: Когда мы создаем программы, нам нужно хранить данные и извлекать данные из памяти. Память не является бесконечным ресурсом. Поэтому разработчик должен следить за тем, чтобы неиспользуемые элементы, хранящиеся в памяти, время от времени уничтожались. Помещение некоторых данных в память называется выделением; обратное действие, которое заключается в удалении данных из памяти, называется освобождением. Роль сборщика мусора заключается в освобождении памяти, когда она больше не используется. Когда в языке нет сборщика мусора, разработчику приходится собирать свой мусор и освобождать память, которая больше не используется... Случайно в Go есть сборщик мусора.

### 6. Состояние Go
Проект рос очень быстро, и сейчас насчитывает более одной тысячи вкладчиков.

На момент написания статьи (8 января 2020 года) последней версией Go является 1.15.6.

Для объединения сообщества организуется множество митапов и конференций.

- В 2018 году было организовано 19 конференций: 3 в США и 16 в других странах.
- В 2017 году было организовано 13 конференций.

Go хотят разработчики:

- В 2018, 2019 и 2020 годах в опросе разработчиков Stackoverflow Go входит в тройку самых востребованных языков программирования [@dev-survy-2018].

### 7. Проверьте себя

#### 7.1 Вопросы
- Какова дата рождения языка Go?
- Что означает параллелизм?
- В среднем, Go показывает очень долгое время сборки? Правда или ложь?

#### 7.2 Ответы
- **Какова дата рождения языка Go?**
  - 2007
- **Что означает параллелизм?**
  - Программа является параллельной, когда задачи могут выполняться одновременно.
- **В среднем, Go показывает очень долгое время сборки? Правда или ложь?**
  - Ложный. Язык был создан именно для решения этой проблемы.

× Бумажное и цифровое издание этой книги доступны [здесь](#).
Я также снял видеокурс по созданию реального проекта с помощью Go.

### 8. Ключевые выводы
- Go родился в 2007 году.
- Go версия 1 была выпущена в 2012 году.
- Язык прост для понимания. Его семантика остается простой.
- Он статически типизирован.
- Он скомпилирован.
- Вы можете писать параллельные программы с помощью Go.
- Это никнейм программистов.

### Библиография
[@go-at-google] ———. 2012. «Go at Google: языковой дизайн на службе программной инженерии». [https://talks.golang.org/2012/splash.article](https://talks.golang.org/2012/splash.article).
[@dev-survy-2018] ———. 2018. «Stack Overflow Developer Survey 2018». [https://insights.stackoverflow.com/survey/2018](https://insights.stackoverflow.com/survey/2018).
[@pike2009go] ———. 2009. «The Go Programming Language». [https://golang.org](https://golang.org).
