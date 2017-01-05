<template>

  <ul class="course-list">
    <li v-for="course in courses" @click="selectCourse(course)">
      {{course.subject}} {{course.number}} - {{course.name}}
    </li>
  </ul>

</template>

<script>

  export default {

    name: 'course-list',

    props: ['courses'],

    methods: {

      selectCourse(course) {

        this.$emit('select-course', course);

      }

    }

  }

</script>

<style>

  .course-list {
    max-height: 500px;
  }

</style>
