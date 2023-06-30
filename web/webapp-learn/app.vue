<template>
  <div class="container-fluid">
    <div class="row">
      <div class="col-4">
        <input v-model="new_kanban.title" placeholder="title" class="form-control">
        <input v-model="new_kanban.body" placeholder="body" class="form-control">
        <button @click="createKanban()" class="btn btn-success">カンバン作成</button>
      </div>
    </div>

    <hr>

    <div class="row mt-3 px-5 h-auto">
      <div v-for="kanban_category in kanban_list" class="col-4">
        <h2 v-if="kanban_list.indexOf(kanban_category) == 0">未着手: {{ kanban_category.length }}</h2>
        <h2 v-else-if="kanban_list.indexOf(kanban_category) == 1">作業中: {{ kanban_category.length }}</h2>
        <h2 v-else-if="kanban_list.indexOf(kanban_category) == 2">完了: {{ kanban_category.length }}</h2>

        <TransitionGroup name="list" tag="ul" class="position-relative list-unstyled">
          <li v-for="k in kanban_category" :key="k" class="card w-100 mx-auto mb-3">
            <div class="card-body">
              <h5 class="card-title">{{ k.title }}</h5>
              <p class="card-text">{{ k.body }}</p>
              <div class="d-flex justify-content-between">
                <button v-if="k.category >= 1" @click="moveKanban(k, 'left')" class="btn btn-dark">◀</button>
                <button v-else class="btn btn-secondary" disabled>◀</button>

                <button v-if="k.category <= 1" @click="moveKanban(k, 'right')" class="btn btn-dark">▶</button>
                <button v-else class="btn btn-secondary" disabled>▶</button>
              </div>
            </div>
          </li>
        </TransitionGroup>
      </div>
    </div>
  </div>
</template>
<script setup>
  class Kanban {
    constructor(title='', body='', category=0) {
      this.title = title
      this.body = body
      this.category = category
    }
  }

  const new_kanban = ref(new Kanban())

  const kanban_list = ref([
    [new Kanban('sample_waiting', 'sample_body')],
    [new Kanban('sample_working', 'sample_body', 1)],
    [new Kanban('sample_completed', 'sample_body', 2)]
  ])

  const createKanban = () => {
    const k = new_kanban.value
    if (k.title == '' || k.body =='') {
      return
    }
    kanban_list.value[0].push(k)
    new_kanban.value = new Kanban()
  }

  const moveKanban = (k, direction) => {
    // kanban_list内の対応するオブジェクトを取得
    const idx = kanban_list.value[k.category].indexOf(k)
    const kanban = kanban_list.value[k.category][idx]

    // 元居たカテゴリから削除
    kanban_list.value[kanban.category].splice(idx, 1)

    // カテゴリ移動
    if (direction == 'left') {
      kanban.category--
    } else {
      kanban.category++
    }
    kanban_list.value[kanban.category].push(kanban)
  }
</script>
<style>
  .list-move,
  .list-enter-active,
  .list-leave-active {
    transition: all 0.5s ease;
  }

  .list-enter-from,
  .list-leave-to {
    opacity: 0;
    transform: translateY(30px);
  }

  .list-leave-active {
    position: absolute;
  }
</style>