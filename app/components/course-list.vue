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

    flex: 1;
    max-height: 100%;
    padding: 0;
    margin: 0;
    border: 1px solid #ddd;
    overflow-y: scroll;

  }

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

</style>
