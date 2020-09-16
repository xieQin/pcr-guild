<template>
<div>
  <van-steps :active="active">
    <van-step>添加队伍</van-step>
    <van-step>添加角色</van-step>
    <van-step>确认阵容</van-step>
  </van-steps>
  <div class="wrapper" v-if="active === 0">
    <van-form @submit="onSubmit" v-if="active === 0">
      <van-field
        v-model="team_name"
        name="阵容名"
        label="阵容名"
        placeholder="阵容名"
        :rules="[{ required: false, message: '请填写阵容名' }]"
      />
      <van-field name="boss_stage" label="boss阶段">
        <template #input>
          <van-radio-group v-model="boss_stage" direction="horizontal">
            <van-radio :name="1">一阶段</van-radio>
            <van-radio :name="2">二阶段</van-radio>
          </van-radio-group>
        </template>
      </van-field>
      <van-field name="boss_num" label="boss">
        <template #input>
          <van-stepper v-model="boss_num" integer min="1" max="6"/>
        </template>
      </van-field>
      <van-field
        v-model="target_damage"
        type="number"
        name="目标伤害"
        label="目标伤害"
        placeholder="目标伤害"
        :rules="[{ required: false, message: '请填写目标伤害' }]"
      />
      <van-field
        v-model="context"
        name="阵容说明"
        label="阵容说明"
        placeholder="阵容说明"
        :rules="[{ required: false, message: '请填写阵容说明' }]"
      />
      <div style="margin-top: 16px;">
        <van-button block type="primary" native-type="submit">
          下一步
        </van-button>
      </div>
    </van-form>
  </div>
  <div class="wrapper" v-if="active === 1">
    <div style="margin: 16px 0;">
      <van-row type="flex" justify="space-between">
        <van-col span="4" v-for="i in 4" :key="i" @click="addTeamCharacter(i)">
          <div class="add-team-character">
            +
          </div>
        </van-col>
        <van-col span="4" @click="addTeamCharacter(5)">
          <div class="add-team-character">
            <img width="50" height="50" :src="`https://pcredivewiki.tw/static/images/unit/icon_unit_104331.png`" alt="">
          </div>
        </van-col>
      </van-row>
    </div>
    <van-form v-if="is_add_character">
      <van-field
        readonly
        clickable
        v-model="unit_name"
        name="picker"
        label="角色"
        placeholder="选择角色"
        @click="showPicker = true"
      />
      <van-popup v-model:show="showPicker" position="bottom">
        <van-picker
          :show-toolbar="true"
          cancel-button-text="取消"
          title="请选择角色"
          :columns="columns"
          @confirm="onConfirm"
          @change="onChange"
          @cancel="showPicker = false"
        />
      </van-popup>
      <van-field
        v-model="level"
        name="等级"
        type="number"
        label="等级"
        placeholder="角色等级"
        :rules="[{ required: false, message: '请填写等级' }]"
      />
      <van-field
        v-model="rarity"
        name="星级"
        type="number"
        label="星级"
        placeholder="角色星级"
        :rules="[{ required: false, message: '请填写星级' }]"
      />
      <van-field
        v-model="promotion"
        name="Rank"
        type="number"
        label="Rank"
        placeholder="角色Rank"
        :rules="[{ required: false, message: '请填写Rank' }]"
      />
      <van-field
        v-model="love_level"
        name="LoveLevel"
        type="number"
        label="好感度"
        placeholder="角色好感度"
        :rules="[{ required: false, message: '请填写好感度' }]"
      />
      <van-field
        v-model="unique_equip_rank"
        name="unique_equip_rank"
        type="number"
        label="专武等级"
        placeholder="角色专武等级"
        :rules="[{ required: false, message: '请填写专武等级' }]"
      />
    </van-form>
    <div style="margin-top: 16px;">
      <van-row type="flex" justify="space-between">
        <van-col span="11">
          <van-button block type="default" native-type="submit" @click="prev()">
            上一步
          </van-button>
        </van-col>
        <van-col span="11">
          <van-button block type="primary" native-type="submit" @click="next()">
            下一步
          </van-button>
        </van-col>
      </van-row>
    </div>
  </div>
  <div class="wrapper" v-if="active === 2">
    <div style="margin-top: 16px;">
      <van-row type="flex" justify="space-between">
        <van-col span="11">
          <van-button block type="default" native-type="submit" @click="prev()">
            上一步
          </van-button>
        </van-col>
        <van-col span="11">
          <van-button to="/team" block type="primary" native-type="submit" @click="confirm()">
            提交
          </van-button>
        </van-col>
      </van-row>
    </div>
  </div>
</div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { Form, Field, RadioGroup, Radio, Stepper, Step, Steps, Col, Row, Picker, Popup } from 'vant'
import { MockCharacters } from '../mock/characters'

@Component({
  components: { Form, Field, RadioGroup, Radio, Stepper, Step, Steps, Col, Row, Picker, Popup }
})
export default class AddTeam extends Vue {
  showPicker = false
  active = 0
  team_name = ""
  target_damage = ""
  boss_num = ""
  boss_stage = 1
  context = ""
  team_characters = []
  columns: string[] = []

  is_add_character = false

  unique_equip_rank = ""
  love_level = ""
  unit_name = ""
  level = ""
  rarity = ""
  promotion = ""

  mounted () {
    const characters: string[] = []
    MockCharacters.forEach(item => characters.push(item.unit_name))
    this.columns = characters || []
  }

  onConfirm (v: any) {
    this.unit_name = v
    this.showPicker = false
  }
  onChange () {}

  addTeamCharacter(i: number) {
    this.is_add_character = !this.is_add_character
  }

  onSubmit() {
    this.is_add_character = false
    this.active ++
  }
  prev() {
    this.is_add_character = false
    this.active --
  }
  next() {
    this.is_add_character = false
    this.active ++
  }
  confirm() {}
}
</script>

<style lang="less">
.add-team-character {
  width: 50px;
  height: 50px;
  border: 1px dashed rgba(69, 90, 100, 0.6);
  color: rgba(69, 90, 100, 0.6);;
  border-radius: 3px;
  text-align: center;
  line-height: 50px;
}
</style>