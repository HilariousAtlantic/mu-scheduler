<template>

  <div class="course-list">

    <header>{{header}}</header>

    <ul>

      <li v-for="course in courses" @click="onCourseClick(course)">

        {{course.subject}} {{course.number}} - {{course.name}}
        <i class="course-info fa fa-info-circle"></i>

      </li>

    </ul>

  </div>



</template>

<script>

  export default {

    name: 'course-list',

    props: ['header', 'courses', 'dispatch'],

    methods: {

      onCourseClick(course) {

        this.$store.dispatch(this.dispatch, course);

      }

    }

  }

</script>

<style lang="postcss" scoped>

  .course-list {

    display: flex;
    flex-direction: column;

  }

  header {

    background: #eee;
    border: 1px solid #ddd;
    padding: 10px;

  }

  ul {

    padding: 0;
    margin: 0;
    flex: 1;
    overflow-y: auto;
    border: 1px solid #ddd;
    border-top: none;

  }

  li {

    list-style: none;
    padding: 10px;
    cursor: pointer;

    &:hover {

      background: #f5f5f5;

      .course-info {

        display: inline-block;

      }

    }

  }

  .course-info {

    display: none;
    float: right;

  }

</style>
