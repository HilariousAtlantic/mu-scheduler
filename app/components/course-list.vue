<template>

  <ul class="course-list">
    <li v-for="course in courses" @click="$emit('select', course)">{{course}}</li>
  </ul>

</template>

<script>

  export default {

    name: 'course-list',

    props: ['courses']

  }

</script>

<style>

  .course-list {
    max-height: 500px;
  }

</style>
