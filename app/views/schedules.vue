<template>

  <div v-if="length" class="schedules-view">

    <double-header>

      <span slot="left">Filtered Schedules</span>

      <span slot="right">Schedule {{index+1}} of {{length}}</span>

    </double-header>

    <schedule-list :index="index" @next="increaseIndex" @prev="decreaseIndex"></schedule-list>

    <router-link to="filters">

      <button class="schedules-button">Change Filters</button>

    </router-link>

  </div>

  <div v-else-if="!$store.state.selectedCourses.length">

    <instruction>

      <div>

        <h3>Viewing Schedules</h3>

        <p>You have not selected any courses.</p>

      </div>

    </instruction>

    <router-link to="courses">

      <button class="schedules-button">Select Courses</button>

    </router-link>

  </div>

  <div v-else-if="!$store.getters.filteredSchedules.length">

    <instruction>

      <div>

        <h3>Viewing Schedules</h3>

        <p>Your filters are eliminating every possible schedule. You will need to edit or remove some in order to find a schedule.</p>

      </div>

    </instruction>

    <router-link to="filters">

      <button class="schedules-button">Change Filters</button>

    </router-link>

  </div>

</template>

<script>

  import Instruction from '../components/common/instruction.vue';
  import DoubleHeader from '../components/common/double-header.vue';
  import ScheduleList from '../components/schedule/schedule-list.vue';

  export default {

    name: 'schedules-view',

    components: {Instruction, DoubleHeader, ScheduleList},

    data() {

      return {

        index: 0

      }

    },

    computed: {

      length() {

        return this.$store.getters.filteredSchedules.length;

      }

    },

    methods: {

      increaseIndex() {

        if (this.index < this.length-1) {

          this.index++;

        }

      },

      decreaseIndex() {

        if (this.index > 0) {

          this.index--;

        }

      }

    }

  }

</script>

<style scoped>

  .instruction {

    margin-bottom: 10px;

    h3 {

      margin-top: 0;

    }

    p {

      margin-top: 0;
      margin-bottom: 10px;

      &:last-of-type {

        margin-bottom: 0;

      }

    }

  }

  .double-header {

    font-weight: 900;
    font-size: 1.25rem;

  }

  .schedule-list {

    margin: 20px 0;

  }

  .schedules-button {

    width: 100%;
    padding: 15px;
    border: 1px solid #ddd;
    background: #4CAF50;
    color: #fff;
    outline: none;
    cursor: pointer;

    &:hover {

      background: #43A047;

    }

  }

</style>
