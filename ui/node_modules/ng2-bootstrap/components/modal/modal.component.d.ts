import { AfterViewInit, ElementRef, EventEmitter, Injector, OnDestroy, Renderer } from '@angular/core';
import { ComponentsHelper } from '../utils/components-helper.service';
import { ModalOptions } from './modal-options.class';
export declare class ModalDirective implements AfterViewInit, OnDestroy {
    private element;
    private renderer;
    private injector;
    private componentsHelper;
    config: ModalOptions;
    onShow: EventEmitter<ModalDirective>;
    onShown: EventEmitter<ModalDirective>;
    onHide: EventEmitter<ModalDirective>;
    onHidden: EventEmitter<ModalDirective>;
    isAnimated: boolean;
    isShown: boolean;
    protected _dialog: any;
    protected _config: ModalOptions;
    protected _isShown: boolean;
    private isBodyOverflowing;
    private originalBodyPadding;
    private scrollbarWidth;
    private backdrop;
    private document;
    /** Host element manipulations */
    protected onClick(event: any): void;
    protected onEsc(): void;
    constructor(element: ElementRef, renderer: Renderer, injector: Injector, componentsHelper: ComponentsHelper);
    ngOnDestroy(): any;
    ngAfterViewInit(): any;
    /** Public methods */
    toggle(): void;
    show(): void;
    hide(event?: Event): void;
    /** Private methods */
    private getConfig(config?);
    /**
     *  Show dialog
     */
    private showElement();
    private hideModal();
    private showBackdrop(callback?);
    private removeBackdrop();
    /** Events tricks */
    private resetAdjustments();
    /** Scroll bar tricks */
    private checkScrollbar();
    private setScrollbar();
    private resetScrollbar();
    private getScrollbarWidth();
}
