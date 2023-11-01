<template>
    <div>
        <h1>TODO List</h1>
        <input type="text" placeholder="Введите название элемента TODO" v-model="title">
        <input type="text" placeholder="Введите текст элемента TODO" v-model="text">
        <br>
        <button @click="pushOne">Добавить в список</button>
        <button @click="getAll">Обновить список</button>

        <li v-for="element, index in todoList" :key="index">
            <p>{{ element.title }}</p>
            <p>{{ element.text }}</p>
        </li>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'TodoList',
    created() { this.getAll() },
    methods: {
        getAll() {
            axios.get('http://localhost:4321/get_all')
                .then((response) => {
                    console.log('response', response)
                    console.log('response.data', response.data)
                    this.todoList = response.data
                })
                .catch((error) => {
                    console.log(error);
                })
        },
        pushOne() {
            if (this.title == '' || this.text == '') { return }

            const newElement = JSON.stringify({
                title: this.title,
                text: this.text,
            })

            console.log('newElement', newElement)

            axios.post('http://localhost:4321/put_one', newElement)
                .then(() => {
                    this.getAll()
                }).catch((error) => {
                    console.log('error', error);
                })
        },
    },
    data() {
        return {
            title: '',
            text: '',
            todoList: []
        }
    }
}
</script>
