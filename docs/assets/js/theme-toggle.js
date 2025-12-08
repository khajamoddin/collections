(function(){
  function applyTheme(t){
    document.documentElement.setAttribute('data-theme', t);
    try{ localStorage.setItem('docs_theme', t); }catch(e){}
  }
  function currentTheme(){
    try{ return localStorage.getItem('docs_theme') || 'light'; }catch(e){ return 'light'; }
  }
  function injectButton(){
    var b = document.createElement('button');
    b.id = 'theme-toggle';
    b.style.position = 'fixed';
    b.style.right = '16px';
    b.style.bottom = '16px';
    b.style.zIndex = '9999';
    b.style.padding = '8px 12px';
    b.style.borderRadius = '8px';
    b.style.border = '1px solid #ddd';
    b.style.background = '#fff';
    b.style.cursor = 'pointer';
    var t = currentTheme();
    b.textContent = t === 'dark' ? 'Light Mode' : 'Dark Mode';
    b.addEventListener('click', function(){
      var nt = currentTheme() === 'dark' ? 'light' : 'dark';
      applyTheme(nt);
      b.textContent = nt === 'dark' ? 'Light Mode' : 'Dark Mode';
    });
    document.body.appendChild(b);
  }
  document.addEventListener('DOMContentLoaded', function(){
    applyTheme(currentTheme());
    injectButton();
  });
})();
