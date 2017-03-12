<template>

  <div class="course">

    <div class="course-info">

      <h2>{{course.subject}} {{course.number}}</h2>
      <p>{{course.title}}</p>
      <span><i class="fa fa-star"></i> {{course.credits}} Credits</span>
      <span><i class="fa fa-circle"></i> {{course.sections.length}} Sections</span>

    </div>

    <button class="btn" @click="$store.dispatch('deselectCourse', course)">Remove</button>

  </div>

</template>

<script>

  export default {

    name: 'course',

    props: ['course'],

    methods: {

      deselectCourse() {

        this.$store.dispatch('deselectCourse', this.course);

      }

    }

  }

</script>

<style scoped>

  .course {
    display: flex;
    align-items: flex-start;
    border-top: 1px #e1e4e8 solid;
    padding: 15px 10px;
  }

  .fa {
    margin-right: 3px;
  }

  .course-info {

    flex: 1;

    span + span {
      margin-left: 10px;
    }

  }

  .course:last-of-type {
    border-bottom: 1px #e1e4e8 solid;
  }

</style>
