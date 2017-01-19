<template>

  <div class="filter-editor">

    <div class="editor-header">

      <span>{{result}}</span>

    </div>

    <div class="editor-section">

      <slot></slot>

    </div>


    <div class="editor-footer">

      <button class="editor-button">Cancel</button>

      <button class="editor-button filter-save">Save</button>

    </div>

  </div>

</template>

<script>

  export default {

    name: 'filter-editor',

    props: ['result']

  }

</script>

<style scoped>

  .filter-editor {

    background: #fff;
    border: 1px solid #ddd;

  }

  .editor-header {

    padding: 10px;
    border-bottom: 1px solid #ddd;
    background: #eee;

  }

  .editor-section {

    padding: 25px;

  }

  .editor-footer {

    display: flex;
    padding: 10px;
    justify-content: flex-end;
    border-top: 1px solid #ddd;

  }

  .editor-button {

    margin-left: 10px;
    padding: 10px 50px;
    border: 1px solid #ddd;
    background: #fff;
    outline: none;
    cursor: pointer;

  }

  .filter-save {

    background: #4CAF50;
    color: #fff;

  }

</style>
