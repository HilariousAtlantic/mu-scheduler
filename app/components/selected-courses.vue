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

<style scoped>

  .selected-courses {

    padding: 0;
    margin: 0;
    border: 1px solid #ddd;
    overflow-y: scroll;
    max-height: 545px;

    li {

      list-style: none;
      padding: 10px;
      border-top: 1px solid #ddd;
      cursor: pointer;

      &:hover {

        background: #f5f5f5;

      }

      &:first-of-type {

        background: #eee;
        border: none;
        cursor: default;

      }

    }

  }

</style>
