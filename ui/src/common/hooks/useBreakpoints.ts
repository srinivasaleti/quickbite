import { useState, useEffect } from "react";

export type Breakpoint = "xs" | "sm" | "m" | "l" | "xl";

export function useBreakpoint(): Breakpoint {
  const getBreakpoint = (width: number): Breakpoint => {
    if (width < 576) return "xs";
    if (width >= 576 && width < 768) return "sm";
    if (width >= 768 && width < 1050) return "m";
    if (width >= 1050 && width < 1440) return "l";
    return "xl";
  };

  const [breakpoint, setBreakpoint] = useState<Breakpoint>(
    getBreakpoint(window.innerWidth)
  );

  useEffect(() => {
    const handleResize = () => {
      setBreakpoint(getBreakpoint(window.innerWidth));
    };

    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  }, []);

  return breakpoint;
}
