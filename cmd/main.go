package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa -framework ApplicationServices -framework Carbon
#include <Cocoa/Cocoa.h>
#include <ApplicationServices/ApplicationServices.h>
#include <Carbon/Carbon.h>
#include <stdlib.h>
#include <stdbool.h>
#include <dispatch/dispatch.h>

// キーコードの定義
#define LEFT_COMMAND_KEYCODE 55
#define RIGHT_COMMAND_KEYCODE 54

// 入力ソース切り替え関数
void switchToEnglish() {
    TISInputSourceRef source = TISCopyInputSourceForLanguage(CFSTR("en-US"));
    if (source) {
        TISSelectInputSource(source);
        CFRelease(source);
    }
}

void switchToJapanese() {
    TISInputSourceRef source = TISCopyInputSourceForLanguage(CFSTR("ja-JP"));
    if (source) {
        TISSelectInputSource(source);
        CFRelease(source);
    }
}

// グローバル変数でCommandキーの状態を管理
static bool leftCommandPressed = false;
static bool rightCommandPressed = false;

// タイマーのディスパッチキュー
static dispatch_queue_t timerQueue;

// タイマーの期間（ナノ秒単位）
#define TIMER_DURATION 50 * NSEC_PER_MSEC // 100ミリ秒

// タイマー設定関数
void setupTimers() {
    timerQueue = dispatch_queue_create("com.example.keytimer", DISPATCH_QUEUE_SERIAL);
}

// 左Commandキーのタイマー開始関数
void startLeftCommandTimer() {
    dispatch_after(dispatch_time(DISPATCH_TIME_NOW, TIMER_DURATION), timerQueue, ^{
        if (leftCommandPressed) {
            switchToEnglish();
            leftCommandPressed = false;
        }
    });
}

// 右Commandキーのタイマー開始関数
void startRightCommandTimer() {
    dispatch_after(dispatch_time(DISPATCH_TIME_NOW, TIMER_DURATION), timerQueue, ^{
        if (rightCommandPressed) {
            switchToJapanese();
            rightCommandPressed = false;
        }
    });
}

// キーイベントコールバック関数
static CGEventRef keyCallback(CGEventTapProxy proxy, CGEventType type, CGEventRef event, void *refcon) {
    if (type == kCGEventFlagsChanged) {
        CGEventFlags flags = CGEventGetFlags(event);
        CGKeyCode keycode = (CGKeyCode)CGEventGetIntegerValueField(event, kCGKeyboardEventKeycode);

        // 左Commandキーの押下/離脱
        if (keycode == LEFT_COMMAND_KEYCODE) {
            if (!(flags & kCGEventFlagMaskCommand)) {
                // 左Commandキーが押された
                leftCommandPressed = true;
                startLeftCommandTimer();
            } else {
                // 左Commandキーが離された
                leftCommandPressed = false;
            }
        }

        // 右Commandキーの押下/離脱
        if (keycode == RIGHT_COMMAND_KEYCODE) {
            if (!(flags & kCGEventFlagMaskCommand)) {
                // 右Commandキーが押された
                rightCommandPressed = true;
                startRightCommandTimer();
            } else {
                // 右Commandキーが離された
                rightCommandPressed = false;
            }
        }
    } else if (type == kCGEventKeyDown || type == kCGEventKeyUp) {
        CGKeyCode keycode = (CGKeyCode)CGEventGetIntegerValueField(event, kCGKeyboardEventKeycode);

        // Commandキー以外のキーが押下された場合、Commandキーの単独押下フラグをリセット
        if (keycode != LEFT_COMMAND_KEYCODE && keycode != RIGHT_COMMAND_KEYCODE) {
            leftCommandPressed = false;
            rightCommandPressed = false;
        }
    }

    return event;
}

// イベントタップの設定と開始
void startKeyTap() {
    setupTimers();

    // イベントマスクの設定
    CGEventMask mask = CGEventMaskBit(kCGEventFlagsChanged) | CGEventMaskBit(kCGEventKeyDown) | CGEventMaskBit(kCGEventKeyUp);

    // イベントタップの作成
    CFMachPortRef tap = CGEventTapCreate(kCGSessionEventTap,
                                        kCGHeadInsertEventTap,
                                        0,
                                        mask,
                                        keyCallback,
                                        NULL);
    if (!tap) {
        NSLog(@"Failed to create event tap");
        exit(1);
    }

    // ループソースの作成
    CFRunLoopSourceRef runLoopSource = CFMachPortCreateRunLoopSource(kCFAllocatorDefault, tap, 0);
    CFRunLoopAddSource(CFRunLoopGetCurrent(), runLoopSource, kCFRunLoopCommonModes);
    CGEventTapEnable(tap, true);

    // ループの開始
    CFRunLoopRun();
}
*/
import "C"

import (
	"fmt"
)

func main() {
	fmt.Println("Starting key tap...")
	go C.startKeyTap()
	select {}
}
