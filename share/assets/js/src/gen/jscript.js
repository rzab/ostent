define(function(require) {
	var React = require('react');
	return {
		mem_rows:        function(Data, $mem)  { return (React.createElement("tr", {key: "mem-rowby-kind-"+$mem.Kind
  }, React.createElement("td", null
    , $mem.Kind), React.createElement("td", {className: "text-right"
    }, $mem.Free), React.createElement("td", {className: "text-right"
    }, $mem.Used, " ", React.createElement("sup", null
      , React.createElement("span", {className: LabelClassColorPercent($mem.UsePercent)
  }, $mem.UsePercent, "%"))), React.createElement("td", {className: "text-right"
    }, $mem.Total))); },
		mem_table:       function(Data, rows)  { return (React.createElement("table", {className: "table1 stripe-table"
  }, React.createElement("thead", null
    , React.createElement("tr", null
      , React.createElement("th", null
        ), React.createElement("th", {className: "text-right"
        }, "Free"), React.createElement("th", {className: "text-right"
        }, "Used"), React.createElement("th", {className: "text-right"
        }, "Total"))), React.createElement("tbody", null
    , rows))); },

		ifbytes_rows:    function(Data, $if)   { return (React.createElement("tr", {key: "ifbytes-rowby-name-"+$if.Name
  }, React.createElement("td", null
    , React.createElement("input", {id: "if-bytes-name-"+$if.Name, className: "collapse-checkbox", type: "checkbox", "aria-hidden": "true", hidden: true
  }), React.createElement("label", {htmlFor: "if-bytes-name-"+$if.Name, className: "clip", style: {maxWidth: '12ch'}
  }, $if.Name)), React.createElement("td", {className: "text-right"
    }, $if.DeltaIn), React.createElement("td", {className: "text-right"
    }, $if.DeltaOut), React.createElement("td", {className: "text-right"
    }, $if.In), React.createElement("td", {className: "text-right"
    }, $if.Out))); },
		ifbytes_table:   function(Data, rows)  { return (React.createElement("table", {className: "table1 stripe-table"
  }, React.createElement("thead", null
    , React.createElement("tr", null
      , React.createElement("th", null
        , "Interface"), React.createElement("th", {className: "text-right nowrap", title: "BITS per second"
        }, "In", React.createElement("span", {className: "unit"
          }, React.createElement("i", null
            , "b"), "ps")), React.createElement("th", {className: "text-right nowrap", title: "BITS per second"
        }, "Out", React.createElement("span", {className: "unit"
          }, React.createElement("i", null
            , "b"), "ps")), React.createElement("th", {className: "text-right nowrap", title: "total BYTES modulo 4G"
        }, "In", React.createElement("span", {className: "unit"
          }, React.createElement("i", null
            , "B"), "%4G")), React.createElement("th", {className: "text-right nowrap", title: "total BYTES modulo 4G"
        }, "Out", React.createElement("span", {className: "unit"
          }, React.createElement("i", null
            , "B"), "%4G")))), React.createElement("tbody", null
    , rows))); },
		iferrors_rows:   function(Data, $if)   { return (React.createElement("tr", {key: "iferrors-rowby-name-"+$if.Name
  }, React.createElement("td", null
    , React.createElement("input", {id: "if-errors-name-"+$if.Name, className: "collapse-checkbox", type: "checkbox", "aria-hidden": "true", hidden: true
  }), React.createElement("label", {htmlFor: "if-errors-name-"+$if.Name, className: "clip", style: {maxWidth: '12ch'}
  }, $if.Name)), React.createElement("td", {className: "text-right"
    }, $if.DeltaIn), React.createElement("td", {className: "text-right"
    }, $if.DeltaOut), React.createElement("td", {className: "text-right"
    }, $if.In), React.createElement("td", {className: "text-right"
    }, $if.Out))); },
		iferrors_table:  function(Data, rows)  { return (React.createElement("table", {className: "table1 stripe-table"
  }, React.createElement("thead", null
    , React.createElement("tr", null
      , React.createElement("th", null
        , "Interface"), React.createElement("th", {className: "text-right nowrap", title: "per second"
        }, "In ", React.createElement("span", {className: "unit"
          }, "ps")), React.createElement("th", {className: "text-right nowrap", title: "per second"
        }, "Out ", React.createElement("span", {className: "unit"
          }, "ps")), React.createElement("th", {className: "text-right nowrap", title: "modulo 4G"
        }, "In ", React.createElement("span", {className: "unit"
          }, "%4G")), React.createElement("th", {className: "text-right nowrap", title: "modulo 4G"
        }, "Out ", React.createElement("span", {className: "unit"
          }, "%4G")))), React.createElement("tbody", null
    , rows))); },
		ifpackets_rows:  function(Data, $if)   { return (React.createElement("tr", {key: "ifpackets-rowby-name-"+$if.Name
  }, React.createElement("td", null
    , React.createElement("input", {id: "if-packets-name-"+$if.Name, className: "collapse-checkbox", type: "checkbox", "aria-hidden": "true", hidden: true
  }), React.createElement("label", {htmlFor: "if-packets-name-"+$if.Name, className: "clip", style: {maxWidth: '12ch'}
  }, $if.Name)), React.createElement("td", {className: "text-right"
    }, $if.DeltaIn), React.createElement("td", {className: "text-right"
    }, $if.DeltaOut), React.createElement("td", {className: "text-right"
    }, $if.In), React.createElement("td", {className: "text-right"
    }, $if.Out))); },
		ifpackets_table: function(Data, rows)  { return (React.createElement("table", {className: "table1 stripe-table"
  }, React.createElement("thead", null
    , React.createElement("tr", null
      , React.createElement("th", null
        , "Interface"), React.createElement("th", {className: "text-right nowrap", title: "per second"
        }, "In ", React.createElement("span", {className: "unit"
          }, "ps")), React.createElement("th", {className: "text-right nowrap", title: "per second"
        }, "Out ", React.createElement("span", {className: "unit"
          }, "ps")), React.createElement("th", {className: "text-right nowrap", title: "total modulo 4G"
        }, "In ", React.createElement("span", {className: "unit"
          }, "%4G")), React.createElement("th", {className: "text-right nowrap", title: "total modulo 4G"
        }, "Out ", React.createElement("span", {className: "unit"
          }, "%4G")))), React.createElement("tbody", null
    , rows))); },

		cpu_rows:        function(Data, $core) { return (React.createElement("tr", {key: "cpu-rowby-N-"+$core.N
  }, React.createElement("td", {className: "text-right nowrap"
    }, $core.N), React.createElement("td", {className: "text-right"
    }, React.createElement("span", {className: $core.UserClass
      }, $core.User)), React.createElement("td", {className: "text-right"
    }, React.createElement("span", {className: $core.SysClass
      }, $core.Sys)), React.createElement("td", {className: "text-right"
    }, React.createElement("span", {className: $core.IdleClass
      }, $core.Idle)))); },
		cpu_table:       function(Data, rows)  { return (React.createElement("table", {className: "table1 stripe-table"
  }, React.createElement("thead", null
    , React.createElement("tr", null
      , React.createElement("th", null
        ), React.createElement("th", {className: "text-right nowrap"
        }, "User", React.createElement("span", {className: "unit"
          }, "%")), React.createElement("th", {className: "text-right nowrap"
        }, "Sys", React.createElement("span", {className: "unit"
          }, "%")), React.createElement("th", {className: "text-right nowrap"
        }, "Idle", React.createElement("span", {className: "unit"
          }, "%")))), React.createElement("tbody", null
    , rows))); },

		dfbytes_rows:    function(Data, $disk) { return (React.createElement("tr", {key: "dfbytes-rowby-dirname-"+$disk.DirName
  }, React.createElement("td", {className: "nowrap"
    }, React.createElement("input", {id: "df-bytes-devname-"+$disk.DevName, className: "collapse-checkbox", type: "checkbox", "aria-hidden": "true", hidden: true
  }), React.createElement("label", {htmlFor: "df-bytes-devname-"+$disk.DevName, className: "clip", style: {maxWidth: '12ch'}
  }, $disk.DevName)), React.createElement("td", {className: "nowrap"
    }, React.createElement("input", {id: "df-bytes-dirname-"+$disk.DirName, className: "collapse-checkbox", type: "checkbox", "aria-hidden": "true", hidden: true
  }), React.createElement("label", {htmlFor: "df-bytes-dirname-"+$disk.DirName, className: "clip", style: {maxWidth: '6ch'}
  }, $disk.DirName)), React.createElement("td", {className: "text-right"
    }, $disk.Avail), React.createElement("td", {className: "text-right"
    }, $disk.Used, " ", React.createElement("sup", null
      , React.createElement("span", {className: LabelClassColorPercent($disk.UsePercent)
  }, $disk.UsePercent, "%"))), React.createElement("td", {className: "text-right"
    }, $disk.Total))); },
		dfbytes_table:   function(Data, rows)  { return (React.createElement("table", {className: "table1 stripe-table"
  }, React.createElement("thead", null
    , React.createElement("tr", null
      , React.createElement("th", {className: "header"
        }, "           ", React.createElement("a", {href: Data.Links.DFFS.Href, className: Data.Links.DFFS.Class
          }, "Device", React.createElement("span", {className: Data.Links.DFFS.CaretClass
            }))), React.createElement("th", {className: "header"
        }, "           ", React.createElement("a", {href: Data.Links.DFMP.Href, className: Data.Links.DFMP.Class
          }, "Mounted", React.createElement("span", {className: Data.Links.DFMP.CaretClass
            }))), React.createElement("th", {className: "header text-right"
        }, React.createElement("a", {href: Data.Links.DFAVAIL.Href, className: Data.Links.DFAVAIL.Class
          }, "Avail", React.createElement("span", {className: Data.Links.DFAVAIL.CaretClass
            }))), React.createElement("th", {className: "header text-right"
        }, React.createElement("a", {href: Data.Links.DFUSED.Href, className: Data.Links.DFUSED.Class
          }, "Used", React.createElement("span", {className: Data.Links.DFUSED.CaretClass
            }))), React.createElement("th", {className: "header text-right"
        }, React.createElement("a", {href: Data.Links.DFTOTAL.Href, className: Data.Links.DFTOTAL.Class
          }, "Total", React.createElement("span", {className: Data.Links.DFTOTAL.CaretClass
            }))))), React.createElement("tbody", null
    , rows))); },
		dfinodes_rows:   function(Data, $disk) { return (React.createElement("tr", {key: "dfinodes-rowby-dirname-"+$disk.DirName
  }, React.createElement("td", {className: "nowrap"
    }, React.createElement("input", {id: "df-inodes-devname-"+$disk.DevName, className: "collapse-checkbox", type: "checkbox", "aria-hidden": "true", hidden: true
  }), React.createElement("label", {htmlFor: "df-inodes-devname-"+$disk.DevName, className: "clip", style: {maxWidth: '12ch'}
  }, $disk.DevName)), React.createElement("td", {className: "nowrap"
    }, React.createElement("input", {id: "df-inodes-devname-"+$disk.DirName, className: "collapse-checkbox", type: "checkbox", "aria-hidden": "true", hidden: true
  }), React.createElement("label", {htmlFor: "df-inodes-devname-"+$disk.DirName, className: "clip", style: {maxWidth: '6ch'}
  }, $disk.DirName)), React.createElement("td", {className: "text-right"
    }, $disk.Ifree), React.createElement("td", {className: "text-right"
    }, $disk.Iused, " ", React.createElement("sup", null
      , React.createElement("span", {className: LabelClassColorPercent($disk.IusePercent)
  }, $disk.IusePercent, "%"))), React.createElement("td", {className: "text-right"
    }, $disk.Inodes))); },
		dfinodes_table:  function(Data, rows)  { return (React.createElement("table", {className: "table1 stripe-table"
  }, React.createElement("thead", null
    , React.createElement("tr", null
      , React.createElement("th", {className: "header"
        }, "Device"), React.createElement("th", {className: "header"
        }, "Mounted"), React.createElement("th", {className: "header text-right"
        }, "Avail"), React.createElement("th", {className: "header text-right"
        }, "Used"), React.createElement("th", {className: "header text-right"
        }, "Total"))), React.createElement("tbody", null
    , rows))); },

		ps_rows:         function(Data, $proc) { return (React.createElement("tr", {key: "ps-rowby-pid-"+$proc.PIDstring
  }, React.createElement("td", {className: "text-right"
    }, " ", $proc.PID), React.createElement("td", {className: "text-right"
    }, " ", $proc.UID), React.createElement("td", null
    , "            ", React.createElement("input", {id: "psuser-pid-"+$proc.PIDstring, className: "collapse-checkbox", type: "checkbox", "aria-hidden": "true", hidden: true
  }), React.createElement("label", {htmlFor: "psuser-pid-"+$proc.PIDstring, className: "clip", style: {maxWidth: '12ch'}
  }, $proc.User)), React.createElement("td", {className: "text-right"
    }, " ", $proc.Priority), React.createElement("td", {className: "text-right"
    }, " ", $proc.Nice), React.createElement("td", {className: "text-right"
    }, " ", $proc.Size), React.createElement("td", {className: "text-right"
    }, " ", $proc.Resident), React.createElement("td", {className: "text-center"
    }, $proc.Time), React.createElement("td", {className: "nowrap"
    }, "     ", React.createElement("input", {id: "psname-pid-"+$proc.PIDstring, className: "collapse-checkbox", type: "checkbox", "aria-hidden": "true", hidden: true
  }), React.createElement("label", {htmlFor: "psname-pid-"+$proc.PIDstring, className: "clip", style: {maxWidth: '42ch'}
  }, $proc.Name)))); },
		ps_table:        function(Data, rows)  { return (React.createElement("table", {className: "table2 stripe-table"
  }, React.createElement("thead", null
    , React.createElement("tr", null
      , React.createElement("th", {className: "header text-right"
        }, " ", React.createElement("a", {href: Data.Links.PSPID.Href, className: Data.Links.PSPID.Class
          }, "PID", React.createElement("span", {className: Data.Links.PSPID.CaretClass
            }))), React.createElement("th", {className: "header text-right"
        }, " ", React.createElement("a", {href: Data.Links.PSUID.Href, className: Data.Links.PSUID.Class
          }, "UID", React.createElement("span", {className: Data.Links.PSUID.CaretClass
            }))), React.createElement("th", {className: "header"
        }, "            ", React.createElement("a", {href: Data.Links.PSUSER.Href, className: Data.Links.PSUSER.Class
          }, "USER", React.createElement("span", {className: Data.Links.PSUSER.CaretClass
            }))), React.createElement("th", {className: "header text-right"
        }, " ", React.createElement("a", {href: Data.Links.PSPRI.Href, className: Data.Links.PSPRI.Class
          }, "PR", React.createElement("span", {className: Data.Links.PSPRI.CaretClass
            }))), React.createElement("th", {className: "header text-right"
        }, " ", React.createElement("a", {href: Data.Links.PSNICE.Href, className: Data.Links.PSNICE.Class
          }, "NI", React.createElement("span", {className: Data.Links.PSNICE.CaretClass
            }))), React.createElement("th", {className: "header text-right"
        }, " ", React.createElement("a", {href: Data.Links.PSVIRT.Href, className: Data.Links.PSVIRT.Class
          }, "VIRT", React.createElement("span", {className: Data.Links.PSVIRT.CaretClass
            }))), React.createElement("th", {className: "header text-right"
        }, " ", React.createElement("a", {href: Data.Links.PSRES.Href, className: Data.Links.PSRES.Class
          }, "RES", React.createElement("span", {className: Data.Links.PSRES.CaretClass
            }))), React.createElement("th", {className: "header text-center"
        }, React.createElement("a", {href: Data.Links.PSTIME.Href, className: Data.Links.PSTIME.Class
          }, "TIME", React.createElement("span", {className: Data.Links.PSTIME.CaretClass
            }))), React.createElement("th", {className: "header"
        }, "            ", React.createElement("a", {href: Data.Links.PSNAME.Href, className: Data.Links.PSNAME.Class
          }, "COMMAND", React.createElement("span", {className: Data.Links.PSNAME.CaretClass
            }))))), React.createElement("tbody", null
    , rows))); },

		vagrant_rows:    function(Data, $mach) { return (React.createElement("tr", {key: "vagrant-rowby-uuid-"+$mach.UUID
  }, React.createElement("td", null
    , "       ", React.createElement("input", {id: "vagrant-uuid-"+$mach.UUID, className: "collapse-checkbox", type: "checkbox", "aria-hidden": "true", hidden: true
  }), React.createElement("label", {htmlFor: "vagrant-uuid-"+$mach.UUID, className: "clip", style: {maxWidth: '7ch'}
  }, $mach.UUID)), React.createElement("td", null
    , "       ", $mach.Name), React.createElement("td", null
    , "       ", $mach.Provider), React.createElement("td", null
    , "       ", React.createElement("input", {id: "vagrant-state-"+$mach.UUID, className: "collapse-checkbox", type: "checkbox", "aria-hidden": "true", hidden: true
  }), React.createElement("label", {htmlFor: "vagrant-state-"+$mach.UUID, className: "clip", style: {maxWidth: '8ch'}
  }, $mach.State)), React.createElement("td", null
    , "       ", React.createElement("input", {id: "vagrant-filepath-"+$mach.UUID, className: "collapse-checkbox", type: "checkbox", "aria-hidden": "true", hidden: true
  }), React.createElement("label", {htmlFor: "vagrant-filepath-"+$mach.UUID, className: "clip", style: {maxWidth: '50ch'}
  }, $mach.Vagrantfile_path)))); },
		vagrant_error:   function(Data)        { return (React.createElement("tr", {key: "vgerror"
  }, React.createElement("td", {colspan: "5"
    }, Data.VagrantError))); },
		vagrant_table:   function(Data, rows)  { return (React.createElement("table", {id: "vgtable", className: "table1 stripe-table"
  }, React.createElement("thead", null
    , React.createElement("tr", null
      , React.createElement("th", null
        , "id"), React.createElement("th", null
        , "name"), React.createElement("th", null
        , "provider"), React.createElement("th", null
        , "state"), React.createElement("th", null
        , "directory"))), React.createElement("tbody", null
    , rows))); }
	};
});
