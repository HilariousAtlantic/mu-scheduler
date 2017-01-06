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

<style scoped>

  .course-list {

    padding: 0;
    margin: 0;
    border: 1px solid #ddd;
    overflow-y: scroll;
    max-height: 500px;

    li {

      list-style: none;
      padding: 10px;
      border-top: 1px solid #ddd;
      cursor: pointer;

      &:hover {

        background: #f5f5f5;

      }

      &:first-of-type {

        border: none;

      }

    }

  }

</style>
