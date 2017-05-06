# workStat
- это утилита, позволяющая (через взаимодействие с api VKontakte) получить данные о трудоустройстве после того или иного факультета.

<p><strong>Чтобы запустить проект:</strong></p>
<ol>
<li>Ставим себе <a href="https://golang.org/doc/install">golang</a></li>
<li>Открываем файл <a href="https://github.com/pankova/workStat/blob/master/config/config.go">config/config.go</a>, дополнительные инструкции прописаны в нем</li>
<li>Из консоли запустить go run <a href="https://github.com/pankova/workStat/blob/master/main.go">main.go</a></li>
</ol>

<p><strong>Что дальше:</strong></p>
<ol>
<li>Утилита попросит вас ввести название страны (пока, для простоты, выбор состоит из трех стран, любую новую страну легко добавить через код региона в том же <a href="https://github.com/pankova/workStat/blob/master/config/config.go">config/config.go</a></li>
<br>

![alt text](images/Снимок%20экрана%202017-05-07%20в%200.11.40.png)

<li>Далее название региона, города, университета и факультета</li>
<br>

![alt text](images/%D0%A1%D0%BD%D0%B8%D0%BC%D0%BE%D0%BA%20%D1%8D%D0%BA%D1%80%D0%B0%D0%BD%D0%B0%202017-05-07%20%D0%B2%200.10.22.png)

<li>This is it! Если все удачно, и с данного факультета есть выпускники (каждый запрос ограничен 1000 результатами, поэтому итоговый вывод займет не более стольки строк), будет выдан список вида</li>

![alt text](images/%D0%A1%D0%BD%D0%B8%D0%BC%D0%BE%D0%BA%20%D1%8D%D0%BA%D1%80%D0%B0%D0%BD%D0%B0%202017-05-07%20%D0%B2%200.23.27.png)

</ol>

<p><strong>Планы:</strong></p>
<ul>
<li>Сделать послойную загрузку результатов (например, по годам выпуска или рождения), чтобы получить полный список, а не ограниченную на уровне api часть.</li>
<li>Прикрутить сюда строковый анализ, чтобы вывод был в виде какой-то статистики, а не списковой простыни.</li>
</ul>

Пока процесс (временно) заморожен, но доработка запланирована на осень'17. Если хотите присоединиться, welcome :)<br>
Проект был сделан в рамках <a href="https://vk.com/sunday_go_school">воскресной школы по Go</a>.
