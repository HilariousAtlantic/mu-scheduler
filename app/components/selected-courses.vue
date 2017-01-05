<template>

  <ul class="selected-courses">
    <li>Selected Courses for {{semester}}</li>
    <li v-for="course in courses" @click="$emit('unselect', course)">{{course}}</li>
  </ul>

</template>

<script>

  export default {

    name: 'selected-courses',

    props: ['semester', 'courses'],

    data() {
      
    }

  }

</script>

<style>

</style>
