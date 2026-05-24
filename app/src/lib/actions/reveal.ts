/**
 * Svelte action for scroll-triggered reveal animations.
 * Uses IntersectionObserver to add a `.revealed` class when elements enter the viewport.
 * Respects prefers-reduced-motion by revealing immediately without animation.
 */

export interface RevealOptions {
	/** Threshold ratio (0-1) to trigger reveal. Default: 0.15 */
	threshold?: number;
	/** Root margin for the observer. Default: '0px 0px -60px 0px' */
	rootMargin?: string;
	/** Delay in ms before adding the class (for staggering). Default: 0 */
	delay?: number;
	/** Whether to only reveal once. Default: true */
	once?: boolean;
}

const defaultOptions: RevealOptions = {
	threshold: 0.05,
	rootMargin: '0px 0px -40px 0px',
	delay: 0,
	once: true
};

// Check for reduced motion preference
function prefersReducedMotion(): boolean {
	if (typeof window === 'undefined') return false;
	return window.matchMedia('(prefers-reduced-motion: reduce)').matches;
}

export function reveal(node: HTMLElement, options: RevealOptions = {}) {
	const opts = { ...defaultOptions, ...options };

	// If user prefers reduced motion, reveal immediately
	if (prefersReducedMotion()) {
		node.classList.add('revealed');
		return { destroy() {} };
	}

	// Set the delay as a CSS custom property for staggering
	if (opts.delay) {
		node.style.setProperty('--reveal-delay', `${opts.delay}ms`);
	}

	const observer = new IntersectionObserver(
		(entries) => {
			for (const entry of entries) {
				if (entry.isIntersecting) {
					const delay = opts.delay || 0;
					if (delay > 0) {
						setTimeout(() => {
							node.classList.add('revealed');
						}, delay);
					} else {
						node.classList.add('revealed');
					}

					if (opts.once) {
						observer.unobserve(node);
					}
				} else if (!opts.once) {
					node.classList.remove('revealed');
				}
			}
		},
		{
			threshold: opts.threshold,
			rootMargin: opts.rootMargin
		}
	);

	observer.observe(node);

	return {
		update(newOptions: RevealOptions) {
			Object.assign(opts, newOptions);
			if (opts.delay) {
				node.style.setProperty('--reveal-delay', `${opts.delay}ms`);
			}
		},
		destroy() {
			observer.disconnect();
		}
	};
}

/**
 * Stagger helper: returns delay value for nth item in a list.
 */
export function stagger(index: number, baseDelay = 80, offset = 0): number {
	return offset + index * baseDelay;
}
