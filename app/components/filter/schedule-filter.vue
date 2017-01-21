<template>

  <div class="schedule-filter">

    <div class="filter-header" @click="toggleEditing">

      <span class="filter-text" :class="{disabled: !active}">{{filterText}}</span>

      <button type="button" class="filter-expand">

        <i class="fa" :class="{'fa-chevron-down': !editing, 'fa-chevron-up': editing}"></i>

      </button>

    </div>

    <div v-if="editing" class="filter-editor">

      <div class="filter-options">

        <slot></slot>

      </div>

      <div class="filter-actions">

        <button v-if="active" type="button" class="filter-disable" @click="toggleActive">Disable</button>

        <button v-else type="button" class="filter-enable" @click="toggleActive">Enable</button>

        <button type="button" class="filter-delete">Remove</button>

      </div>

    </div>

  </div>

</template>

<script>

  export default {

    name: 'schedule-filter',

    props: ['id', 'text', 'active'],

    data() {

      return {

        editing: false

      }

    },

    computed: {

      filterText() {

        return this.text + (this.active ? '' : ' (Disabled)');

      }

    },

    methods: {

      toggleEditing() {

        this.editing = !this.editing;

      },

      toggleActive() {

        this.$store.dispatch('toggleScheduleFilter', this.id);

      }

    }

  }

</script>

<style scoped>

  .schedule-filter {

    margin-bottom: 10px;

  }

  .filter-header {

    display: flex;
    align-items: center;
    padding: 10px;
    background: #eee;
    border: 1px solid #ddd;
    cursor: pointer;

  }

  .filter-text {

    flex: 1;

  }

  .filter-text.disabled {

    color: #aaa;

  }

  .filter-expand {

    background: #eee;
    border: none;
    outline: none;

  }

  .filter-editor {

    display: flex;
    padding: 10px;
    border: 1px solid #ddd;
    border-top: none;

  }

  .filter-options {

    flex: 1;

  }

  .filter-delete {

    background: #F44336;
    color: #fff;
    border: 1px solid #ddd;
    padding: 10px 25px;
    outline: none;
    cursor: pointer;

  }

  .filter-enable {

    background: #4CAF50;
    color: #fff;
    border: 1px solid #ddd;
    padding: 10px 25px;
    outline: none;
    cursor: pointer;

  }

  .filter-disable {

    background: #FF9800;
    color: #fff;
    border: 1px solid #ddd;
    padding: 10px 25px;
    outline: none;
    cursor: pointer;

  }

</style>
