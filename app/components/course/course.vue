<template>

  <div class="course">

    <div class="course-head" @click="toggleExpanded">

      <span class="course-name">{{course.subject}} {{course.number}} - {{course.title}}</span>

      <button type="button" class="course-expand">

        <i class="fa" :class="{'fa-chevron-down': !expanded, 'fa-chevron-up': expanded}"></i>

      </button>

    </div>

    <div v-if="expanded" class="course-body">

      <double-header>

        <div slot="left">

          <span>{{course.credits}} Credits</span>

          <span>&middot;</span>

          <span>{{course.sections.length}} Sections</span>

        </div>

        <button class="course-deselect" slot="right" @click="deselectCourse">Remove Course</button>

      </double-header>

    </div>

  </div>

</template>

<script>

  import DoubleHeader from '../common/double-header.vue';

  export default {

    name: 'course',

    props: ['course'],

    components: {DoubleHeader},

    data() {

      return {

        expanded: false

      }

    },

    methods: {

      deselectCourse() {

        this.$store.dispatch('deselectCourse', this.course);

      },

      toggleExpanded() {

        this.expanded = !this.expanded;

      }

    }

  }

</script>

<style scoped>

  .course {

    margin-bottom: 10px;

  }

  .course-head {

    display: flex;
    align-items: center;
    padding: 10px;
    background: #eee;
    border: 1px solid #ddd;
    cursor: pointer;

  }

  .course-name {

    flex: 1;

  }

  .course-expand {

    background: #eee;
    border: none;
    outline: none;

  }

  .course-body {

    padding: 10px;
    border: 1px solid #ddd;
    border-top: none;

  }

  .double-header {

    align-items: center;

  }

  .course-deselect {

    background: #F44336;
    color: #fff;
    border: 1px solid #ddd;
    padding: 10px 25px;
    outline: none;
    cursor: pointer;

  }

</style>
