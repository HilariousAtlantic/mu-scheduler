<template>

  <ul class="course-list">
    <li>{{header}}</li>
    <li v-for="course in courses" @click="selectCourse(course)">
      {{course.subject}} {{course.number}} - {{course.name}}
    </li>
  </ul>

</template>

<script>

  export default {

    name: 'course-list',

    props: ['header', 'courses'],

    methods: {

      selectCourse(course) {

        this.$emit('clickCourse', course);

      }

    }

  }

</script>

<style scoped>

  .course-list {

    padding: 0;
    margin: 0;

    &:first-of-type {

      padding-right: 10px;

    }

  }

  li {

    list-style: none;
    padding: 10px;
    border: 1px solid #ddd;
    border-top: none;
    cursor: pointer;

    &:hover {

      background: #f5f5f5;

    }

    &:first-of-type {

      background: #eee;
      border-top: 1px solid #ddd;
      cursor: default;

    }

  }

</style>
