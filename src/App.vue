<!-- https://www.youtube.com/watch?v=qXhmJhR0Fhg&list=PL-wEcSTifrSn5cae0gFQ7Gy7v3t6c7hLF -->

<template>
  <MyComponent />
  <CounterVue name="Яблоки" :min=1 :max=10 />
</template>

<script>
import MyComponent from './components/MyComponent.vue';
import CounterVue from './components/Counter.vue';

export default {
  name: 'App',
  components: {
    MyComponent,
    CounterVue,
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
