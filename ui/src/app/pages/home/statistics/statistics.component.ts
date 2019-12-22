import { Component, OnInit, ViewChild } from '@angular/core';
import {TaskService} from '../../../services/task.service';
import StatisticsData from '../../../models/data/statisticsData';

import * as echarts from 'echarts';
import ECharts = echarts.ECharts;
import EChartOption = echarts.EChartOption;


@Component({
  selector: 'app-statistics',
  templateUrl: './statistics.component.html',
  styleUrls: ['./statistics.component.scss']
})
export class StatisticsComponent implements OnInit {
  @ViewChild('taskState', { static: false }) taskStateView;
  @ViewChild('commentLine', { static: false }) commentLineView;
  private taskStateOption: EChartOption;
  private taskStateChart: ECharts;
  private statisticsData = new StatisticsData();

  private commentLineOption: EChartOption;
  private commentLineChart: ECharts;


  constructor(private taskService: TaskService) {
  }

  ngOnInit() {
    this.taskService.statistics().then((data) => {
      if (data !== undefined) {
        this.statisticsData = data;
        this.taskStateChart = echarts.init(this.taskStateView.nativeElement, 'macarons');
        this.commentLineChart = echarts.init(this.commentLineView.nativeElement, 'macarons');
        this.initTaskStateOptions();
        this.initcommentLineOptions();
      }
    });
  }

  convertStateNumToStr(num: number): string {
    const mp = {};
    mp[0] = '没有进展';
    mp[1] = '即将开始';
    mp[2] = '进行中';
    mp[3] = '完成';
    mp[4] = '取消';
    mp[5] = '暂停';
    return mp[num];
  }

  initTaskStateOptions() {
    const vthis = this;
    const data = [];
    const titles = [];
    for (const ts of this.statisticsData.tasks) {
      data.push({value: ts.count, name: this.convertStateNumToStr(ts.state)});
      titles.push(this.convertStateNumToStr(ts.state));
    }
    this.taskStateOption = {
      title : {
        text: '任务状态概览',
        subtext: '全部数据',
        x: 'center'
      },
      tooltip : {
        trigger: 'item',
        formatter: '{a} <br/>{b} : {c} ({d}%)'
      },
      legend: {
        orient: 'vertical',
        left: 'left',
        data: titles,
      },
      series : [
        {
          name: '状态',
          type: 'pie',
          radius : '55%',
          center: ['50%', '60%'],
          data,
          itemStyle: {
            emphasis: {
              shadowBlur: 10,
              shadowOffsetX: 0,
              shadowColor: 'rgba(0, 0, 0, 0.5)'
            }
          }
        }
      ]
    } as EChartOption;
    this.taskStateChart.setOption(this.taskStateOption);
  }

  initcommentLineOptions() {
    const vthis = this;
    const data = [];
    const titles = [];
    for (const ts of this.statisticsData.comments) {
      data.push(ts.count);
      titles.push(ts.date);
    }
    this.commentLineOption = {
      title: {

        text: '任务状态更新频率',
        subtext: '最近 30 天',
        x: 'center'
      },
      tooltip: {},
      xAxis: {
        data: titles,
        silent: false,
        splitLine: {
          show: false
        }
      },
      yAxis: {
      },
      series: [{
        name: 'bar',
        type: 'bar',
        data,
        animationDelay(idx) {
          return idx * 10;
        }
      }],
      animationEasing: 'elasticOut',
      animationDelayUpdate(idx) {
        return idx * 5;
      }
    } as EChartOption;
    this.commentLineChart.setOption(this.commentLineOption);

  }

}
