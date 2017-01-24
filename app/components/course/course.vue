<template>

  <div class="course">

    <div class="course-head" @click="toggleExpanded">

      <span class="course-name">{{course.subject}} {{course.number}} - {{course.title}}</span>

      <button type="button" class="course-expand">

        <i class="fa" :class="{'fa-chevron-down': !expanded, 'fa-chevron-up': expanded}"></i>

      </button>

    </div>

    <div v-if="expanded" class="course-body">

      <div class="course-sections">

        <span v-for="section in sections">{{section}}</span>

      </div>

      <div class="course-actions">

        <span>{{course.credits}} Credits</span>

        <button class="course-deselect" type="button" @click="deselectCourse">Remove Course</button>

      </div>

    </div>

  </div>

</template>

<script>

  export default {

    name: 'course',

    props: ['course'],

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

    },

    computed: {

      sections() {

        return this.course.sections.sort((sectionA, sectionB) => {

          return sectionA.name.localeCompare(sectionB.name);

        }).map(({name, meets}) => {

          return 'Section ' + name + ' - ' + meets.map(({instructor}) => instructor).join(',');

        });

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

  .course-body {

    display: flex;
    padding: 10px;
    border: 1px solid #ddd;
    border-top: none;

  }

  .course-sections {

    flex: 1;

    span {

      display: block;

    }

  }

  .course-actions {

    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    span {

      font-weight: 900;
      margin-bottom: 10px;

    }

  }

  .course-name {

    flex: 1;

  }

  .course-expand {

    background: #eee;
    border: none;
    outline: none;

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
