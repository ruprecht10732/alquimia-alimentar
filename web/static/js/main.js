// Lógica do Overlay Menu Mobile
const btnOpen = document.getElementById('mobile-menu-btn');
const btnClose = document.getElementById('close-menu-btn');
const mobileMenu = document.getElementById('mobile-menu');
const mobileLinks = document.querySelectorAll('.mobile-link');

const openMenu = () => {
    mobileMenu.classList.remove('translate-x-full');
    document.body.style.overflow = 'hidden';
};

const closeMenu = () => {
    mobileMenu.classList.add('translate-x-full');
    document.body.style.overflow = '';
};

btnOpen.addEventListener('click', openMenu);
btnClose.addEventListener('click', closeMenu);

mobileLinks.forEach(link => {
    link.addEventListener('click', closeMenu);
});

// Header Sticky
const header = document.getElementById('main-header');
window.addEventListener('scroll', () => {
    if (window.scrollY > 20) {
        header.classList.add('shadow-sm', 'border-matcha-dark/10');
    } else {
        header.classList.remove('shadow-sm', 'border-matcha-dark/10');
    }
});

// Efeito Reveal Profissional (Animação ao scrollar a tela)
const observerOptions = {
    root: null,
    rootMargin: '0px 0px -10% 0px',
    threshold: 0.1
};

const observer = new IntersectionObserver((entries, observer) => {
    entries.forEach(entry => {
        if (entry.isIntersecting) {
            entry.target.classList.add('active');
            observer.unobserve(entry.target);
        }
    });
}, observerOptions);

document.querySelectorAll('.reveal').forEach(el => {
    observer.observe(el);
});
