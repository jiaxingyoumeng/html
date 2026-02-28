const Store = {
  key: 'doremi-biz-v2',
  defaults: {
    todayOrders: 28,
    monthRevenue: 586320,
    newLeads: 45,
    pendingAlbums: 12,
    revenueTarget: 700000,
    orders: [
      { id: 'OD24001', customer: '张先生', pkg: '尊享策划', status: '待签约', amount: 38800 },
      { id: 'OD24002', customer: '李女士', pkg: '摄影套餐', status: '执行中', amount: 16800 },
      { id: 'OD24003', customer: '王先生', pkg: '跟拍套餐', status: '已完成', amount: 9800 }
    ],
    leads: [
      { name: '赵女士', stage: '新线索', next: '今晚电话沟通' },
      { name: '陈先生', stage: '已邀约', next: '周三到店咨询' },
      { name: '刘女士', stage: '已到店', next: '等待报价确认' }
    ],
    albums: [
      { customer: '张先生', name: '草坪婚礼精选', status: '修图中', progress: 45 },
      { customer: '李女士', name: '夜景仪式合集', status: '待客户确认', progress: 90 },
      { customer: '王先生', name: '酒店仪式全片', status: '已交付', progress: 100 }
    ],
    todos: [
      { level: '紧急', text: '确认张先生婚礼方案', done: false },
      { level: '普通', text: '跟进李女士选片', done: false },
      { level: '普通', text: '核对本周排期', done: true }
    ],
    settings: { shop: 'Doremi 婚礼策划中心', phone: '400-800-9000', sms: true, mail: true }
  },
  load() {
    try { return { ...this.defaults, ...(JSON.parse(localStorage.getItem(this.key)) || {}) }; }
    catch { return { ...this.defaults }; }
  },
  save(patch) {
    const next = { ...this.load(), ...patch };
    localStorage.setItem(this.key, JSON.stringify(next));
    window.dispatchEvent(new CustomEvent('store:change', { detail: next }));
    return next;
  }
};

function nav(active) {
  const links = [
    ['index.html', '控制台'],
    ['orders.html', '订单管理'],
    ['customers.html', '客户管理'],
    ['albums.html', '图册交付'],
    ['analytics.html', '经营分析'],
    ['settings.html', '系统设置'],
  ];
  const el = document.querySelector('#menu');
  if (!el) return;
  el.innerHTML = links.map(([href, text]) => `<a class="${active===href?'active':''}" href="${href}"><span>${text}</span><span>›</span></a>`).join('');
}

function notice(msg) {
  const el = document.querySelector('#notice');
  if (!el) return;
  el.textContent = msg;
  el.classList.add('show');
  clearTimeout(window.__nt);
  window.__nt = setTimeout(() => el.classList.remove('show'), 1800);
}

function statusBadge(status){
  if (['已完成','已交付','已到店'].includes(status)) return 'b-success';
  if (['执行中','待客户确认','已邀约'].includes(status)) return 'b-danger';
  return 'b-warn';
}
