# Команда

Команда — это поведенческий паттерн проектирования, который превращает запросы в объекты, позволяя передавать их как
аргументы при вызове методов, ставить запросы в очередь, логировать их, а также поддерживать отмену операций.

Применимость:

Когда мы хотим параметризовать объекты выполняемым действием команда превращает операции в объекты. А объекты можно
передавать, хранить и взаимозаменять внутри других объектов.

Когда мы хотитм ставить операции в очередь, выполнять их по расписанию или передавать по сети.
Как и любые другие объекты, команды можно сериализовать, то есть превратить в строку, чтобы потом сохранить в файл или
базу данных. Затем в любой удобный момент её можно достать обратно, снова превратить в объект команды и выполнить. Таким
же образом команды можно передавать по сети, логировать или выполнять на удалённом сервере.

Когда нам нужна операция отмены.
Главная вещь, которая вам нужна, чтобы иметь возможность отмены операций, — это хранение истории. Среди многих способов,
которыми можно это сделать, паттерн Команда является, пожалуй, самым популярным.

Плюсы + :

1) Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют;
2) Позволяет реализовать простую отмену и повтор операций;
3) Позволяет реализовать отложенный запуск операций;
4) Позволяет собирать сложные команды из простых;
5) Реализует принцип открытости/закрытости.

Минусы - :

1) Усложняет код программы из-за введения множества дополнительных классов.

Пример из практики:

1) Графические интерфейсы (GUI):
   В системах GUI команды могут быть представлены различными действиями, такими как нажатие кнопки, выбор опции в меню и
   т.д.
   Команды могут использоваться для реализации отмены и повтора операций.
2) Системы управления транзакциями:
   Команды могут использоваться для представления транзакций в базах данных.
   Каждая команда может представлять собой отдельную операцию, которая может быть отменена в случае необходимости.
3) Обработка пользовательских запросов:
   В веб-приложениях команды могут использоваться для обработки запросов от клиента.
   Например, команда может быть создана для обработки запроса на добавление нового элемента или изменение настроек.
4) Системы плагинов:
   Команды могут служить основой для систем плагинов, где каждый плагин представляет собой команду.
   Это обеспечивает гибкость в добавлении новых функциональностей без изменения основного кода.
