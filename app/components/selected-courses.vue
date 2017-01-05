<template>

  <ul class="selected-courses">
    <li>Selected Courses for {{semester}}</li>
    <li v-for="course in courses" @click="unselectCourse(course)">
      {{course.subject}} {{course.number}} - {{course.name}}
    </li>
  </ul>

</template>

<script>

  export default {

    name: 'selected-courses',

    props: ['semester', 'courses'],

    methods: {

      unselectCourse(course) {

        this.$emit('unselect-course', course);

      }

    }

  }

</script>

<style>

  .selected-courses li:first-of-type {
    background: #eee;
    cursor: default;
  }

  .selected-courses {
    max-height: 545px;
  }

</style>
